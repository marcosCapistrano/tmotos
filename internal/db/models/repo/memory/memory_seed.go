package memory

func (repo *MemoryRepo) Seed() {
	root := repo.productTree.root

	pecas := root.Insert("Peças", "pecas", false)
	vestuario := root.Insert("Vestuário", "vestuario", false)

	transmissão := pecas.Insert("Transmissão", "transmissao", false)
	capacetes := vestuario.Insert("Capacetes", "capacetes", false)

	transmissão.Insert("Relação", "relacao", true)
	escamote := capacetes.Insert("Escamoteável", "escamoteavel", false)
	fixo := capacetes.Insert("Fixo", "fixo", false)

	escamote.Insert("ASX", "asx-40", true)
	fixo.Insert("LS2", "ls2", true)
}
