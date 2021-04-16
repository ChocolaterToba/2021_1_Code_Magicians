package repository

import "pinterest/domain/entity"

type BoardRepository interface {
	AddBoard(board *entity.Board) (int, error) // Creating user's board
	GetBoard(int) (*entity.Board, error)       // Get description of the board
	GetBoards(int) ([]entity.Board, error)     // Get boards by authorID
	DeleteBoard(int, int) error                // Removes user's board by ID
	GetInitUserBoard(int) (int, error)         // Get initial user's board
	CheckBoard(int, int) error                 // Checking that passed board belongs to passed user
	UploadBoardAvatar(int, string) error
}
