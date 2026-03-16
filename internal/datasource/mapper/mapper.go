package mapper

import (
	datasourceModel "tic-tac-toe/internal/datasource/models"
	domainModel "tic-tac-toe/internal/domain/models"
)

func ToDatasource(domainEntity domainModel.CurrentGame) datasourceModel.CurrentGame {
	return datasourceModel.CurrentGame{
		UUID: domainEntity.UUID,
		Grid: domainEntity.Grid.Matrix,
	}
}

func ToDomain(datasourceEntity datasourceModel.CurrentGame) domainModel.CurrentGame {
	return domainModel.CurrentGame{
		UUID: datasourceEntity.UUID,
		Grid: domainModel.Grid{Matrix: datasourceEntity.Grid},
	}
}
