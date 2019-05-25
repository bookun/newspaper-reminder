package model

import "fmt"

type Person interface {
	GetMessage() string
}

type Man struct {
	ID string
}

func (m *Man) GetMessage() string {
	return fmt.Sprintf("おい <@%s>、新聞持ってこい", m.ID)
}

type Girl struct {
	ID string
}

func (g *Girl) GetMessage() string {
	return fmt.Sprintf("すみません <@%s>さん、新聞を持ってきていただいてもよろしいでしょうか？\n 明日は違うと思います。お手数おかけします。コーヒー淹れて待ってます。", g.ID)
}
