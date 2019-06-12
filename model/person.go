package model

import "fmt"

type Person interface {
	GetMessage() string
	GetID() string
	GetResetMessage() string
}

type Man struct {
	ID string
}

func (m *Man) GetMessage() string {
	return fmt.Sprintf("おい <@%s>、新聞持ってこい", m.ID)
}

func (m *Man) GetResetMessage() string {
	return fmt.Sprintf("<@%s>が `reset` を放った", m.ID)
}

func (m *Man) GetID() string {
	return m.ID
}

type Girl struct {
	ID string
}

func (g *Girl) GetMessage() string {
	return fmt.Sprintf("すみません <@%s>さん、新聞を持ってきていただいてもよろしいでしょうか？\n 明日は違うと思います。お手数おかけします。コーヒー淹れて待ってます。", g.ID)
}

func (g *Girl) GetID() string {
	return g.ID
}

func (g *Girl) GetResetMessage() string {
	return fmt.Sprintf("<@%s>さんが `reset` を打ってくれました。", g.ID)
}
