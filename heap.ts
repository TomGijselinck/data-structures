interface HasWeight {
    weight: number
}

function newPriorityQueue<T>(): T[] {
    return []
}

function min<T extends HasWeight>(pq: T[]): T {
    const minElement = pq[0]
    const last = pq.pop()
    if (pq.length > 0 && last != undefined) {
        pq[0] = last
        sink(pq, 0)
    }
    return minElement
}

function sink<T extends HasWeight>(pq: T[], i: number) {
    const left = 2*i + 1
    const right = left + 1
    let smallest = i
    if (left < pq.length && pq[left].weight < pq[i].weight) {
        smallest = left
    }
    if (right < pq.length && pq[right].weight < pq[smallest].weight) {
        smallest = right
    }
    if (smallest != i) {
        swap(pq, i, smallest)
        sink(pq, smallest)
    }
}

function swap<T>(pq: T[], i: number, j: number) {
    const tmp = pq[i]
    pq[i] = pq[j]
    pq[j] = tmp
}

function add<T extends HasWeight>(pq: T[], element: T) {
    pq.push(element)
    swim(pq, pq.length - 1)
}

function swim<T extends HasWeight>(pq: T[], i: number) {
    const parent = Math.floor((i-1)/2)
    if (i > 0 && pq[parent].weight > pq[i].weight) {
        swap(pq, parent, i)
        swim(pq, parent)
    }
}

function notEmpty<T>(pq: T[]): boolean {
    return pq.length > 0
}