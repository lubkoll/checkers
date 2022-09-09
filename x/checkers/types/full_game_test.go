package types

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/alice/checkers/x/checkers/rules"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func GetStoredGame1() *StoredGame {
	return &StoredGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Index:   "1",
		Game:    rules.New().String(),
		Turn:    "b",
	}
}

func TestCanGetAddressCreator(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	creatorAddress, err2:= GetStoredGame1().GetCreatorAddress()
	require.Equal(t, aliceAddress, creatorAddress)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongCreator(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Creator = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4"
	creatorAddress, err := storedGame.GetCreatorAddress()
	require.Nil(t, creatorAddress)

	require.EqualError(t,
		err,
		"creator address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: checksum failed. Expected 3xn9d3, got 3xn9d4.")
	// Validate() does not exist for StoredGame
	// require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestCanGetRedPlayer(t *testing.T) {
	carolAddress, err1 := sdk.AccAddressFromBech32(carol)
	redAddress, err2:= GetStoredGame1().GetRedAddress()
	require.Equal(t, carolAddress, redAddress)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongRedPlayer(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Red = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd8"
	redAddress, err := storedGame.GetRedAddress()
	require.Nil(t, redAddress)

	require.EqualError(t,
		err,
		"red address is invalid: cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd8: decoding bech32 failed: checksum failed. Expected 4yuxd7, got 4yuxd8.")
	// Validate() does not exist for StoredGame
	// require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestCanGetBlackPlayer(t *testing.T) {
	bobAddress, err1 := sdk.AccAddressFromBech32(bob)
	blackAddress, err2:= GetStoredGame1().GetBlackAddress()
	require.Equal(t, bobAddress, blackAddress)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongBlackPlayer(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h"
	blackAddress, err := storedGame.GetBlackAddress()
	require.Nil(t, blackAddress)

	require.EqualError(t,
		err,
		"black address is invalid: cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h: decoding bech32 failed: checksum failed. Expected xqhc8g, got xqhc8h.")
	// Validate() does not exist for StoredGame
	// require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestParseGameCorrect(t *testing.T) {
	game, err := GetStoredGame1().ParseGame()
	require.EqualValues(t, rules.New().Pieces, game.Pieces)
	require.Nil(t, err)
}

func TestParseGameCanIfChangedOk(t *testing.T) {
	storedGame := GetStoredGame1()
	// t.Error(storedGame.Game)
	storedGame.Game = strings.Replace(storedGame.Game, "b", "r", 1)
	t.Error(storedGame.Game)
	game, err := storedGame.ParseGame()
	require.NotEqualValues(t, rules.New().Pieces, game)
	require.Nil(t, err)
}
