package main

import  "github.com/"

func init() {
	log.SetFlags(log.Lshortfile)
	cmd.EnsureAgendaDir()
}

func main() {
	cmd.Execute()
}