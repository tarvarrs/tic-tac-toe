package mapper

import (
	dsModel "tic-tac-toe/internal/datasource/models"
	domainModel "tic-tac-toe/internal/domain/models"
)

func ToDatasource(de domainModel.CurrentGame) dsModel.CurrentGame {
	return dsModel.CurrentGame{
		UUID: de.UUID,
		Grid: de.Grid.Matrix,
	}
}

func ToDomain(dse dsModel.CurrentGame) domainModel.CurrentGame {
	return domainModel.CurrentGame{
		UUID: dse.UUID,
		Grid: domainModel.Grid{Matrix: dse.Grid},
	}
}
