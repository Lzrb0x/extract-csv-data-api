package model

type Driver struct {
	ID   uint   `csv:"id"`
	NOME string `csv:"nome"`
	CNH  string `csv:"cnh"`
}
