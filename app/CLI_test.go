package poker_test

import (
	poker "app"
	"bytes"
	"io"
	"strings"
	"testing"
)

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCLI(t *testing.T) {

	t.Run("start game with 3 players and finish the game with 'Chris' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}
		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertGameStartedWith(t, game, 3)
		poker.AssertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, poker.DummyStdOut, game)
		cli.PlayPoker()

		poker.AssertGameStartedWith(t, game, 8)
		poker.AssertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}
		in := userSends("pies")
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}
		in := userSends("8", "Lloyd is a killer")
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMsg)
	})
}

func assertGameNotFinished(t testing.TB, game *poker.GameSpy) {
	t.Helper()
	if game.FinishCalled {
		t.Errorf("game should not have finished")
	}
}

func assertGameNotStarted(t testing.TB, game *poker.GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertMessageSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %q", got, messages)
	}
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got amount %+v want %+v", got, want)
	}
}
