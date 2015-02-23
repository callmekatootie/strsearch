package strsearch

import "testing"

func TestFind(t *testing.T) {
    text := "The ordinance or emergency executive order has to be cleared by both Houses of Parliament within six weeks from today."
    pattern := "Parliament"

    position := Find(text, pattern)
    if position != 79 {
        t.Errorf("Expected %d, got %d", 79, position)
    }

    text = "British actor Eddie Redmayne has won the best actor Oscar for The Theory of Everything, while Julianne Moore picked up best actress for Still Alice."
    pattern = "home"

    position = Find(text, pattern)
    if position != -1 {
        t.Errorf("Expected %d, got %d", -1, position)
    }
}
