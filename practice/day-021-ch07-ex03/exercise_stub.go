package main

import "io"

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {}
