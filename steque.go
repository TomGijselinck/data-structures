package main

type Steque struct {
	start *LinkNode
	end   *LinkNode
	size  int
}

type LinkNode struct {
	item  int
	left  *LinkNode
	right *LinkNode
}

func (s *Steque) pushLeft(value int) {
	oldStart := s.start
	s.start = &LinkNode{value, nil, oldStart}
	if oldStart == nil {
		s.end = s.start
	} else {
		oldStart.left = s.start
	}
	s.size++
}

func (s *Steque) pushRight(value int) {
	oldEnd := s.end
	s.end = &LinkNode{item: value, left: oldEnd}
	if oldEnd == nil {
		s.start = s.end
	} else {
		oldEnd.right = s.end
	}
	s.size++
}

func (s *Steque) popLeft() int {
	if s.size == 0 {
		panic("cannot pop left empty steque")
	}
	oldStart := s.start
	s.start = oldStart.right
	if s.start != nil {
		s.start.left = nil
	}
	oldStart.right = nil
	s.size--
	return oldStart.item
}

func (s *Steque) popRight() int {
	if s.size == 0 {
		panic("cannot pop right empty steque")
	}
	oldEnd := s.end
	s.end = oldEnd.left
	if s.end != nil {
		s.end.right = nil
	}
	oldEnd.left = nil
	s.size--
	return oldEnd.item
}
