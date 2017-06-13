package main

type available struct{
	Name string     `json:"name"`
	Lang string     `json:"lang"`
	Ver  []int   `json:"version"`
}

type Todos []available
