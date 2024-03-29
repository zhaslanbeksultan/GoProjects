package models

import "strings"

type Player struct {
	Name     string `json:"name"`
	Age      string `json:"age"`
	Number   string `json:"number"`
	Position string `json:"position"`
	Image    string `json:"image"`
}

var players = []Player{
	{Name: "Ter Stegen", Age: "30", Number: "1", Position: "Goalkeeper", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/7d67f42b-2a8d-4e51-8d41-8aa3542482e4/01-MARC-ANDRE_TER_STEGEN.jpg?width=470&height=470"},
	{Name: "Inaki Pena", Age: "24", Number: "13", Position: "Goalkeeper", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/65ae7c6c-8b9a-4d94-94b3-273c30c596a5/26-INAKI_PENA.jpg?width=470&height=470"},
	{Name: "Cancelo", Age: "26", Number: "2", Position: "Defender", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/65ae7c6c-8b9a-4d94-94b3-273c30c596a5/26-INAKI_PENA.jpg?width=470&height=470"},
	{Name: "Balde", Age: "19", Number: "3", Position: "Defender", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/993af0d7-2b2c-40d1-9e9b-c24f8d165537/mini_28-BALDE.jpg?width=470&height=470"},
	{Name: "Araujo", Age: "23", Number: "4", Position: "Defender", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/93330ade-5bdf-420a-a989-2c337f8104ca/04-RONALD_ARAUJO_.jpg?width=470&height=470"},
	{Name: "Martinez", Age: "25", Number: "5", Position: "Defender", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/8e1f4a73-cca2-48d8-8e43-655334f27886/jugador_fitxa_inigo_2023.jpg?width=470&height=470"},
	{Name: "Christinsen", Age: "27", Number: "15", Position: "Defender", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/3acd07d7-8081-4050-a62c-1c1302038c23/15-ANDREAS_CHRISTENSEN.jpg?width=470&height=470"},
	{Name: "Alonso", Age: "28", Number: "17", Position: "Defender", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/383058f5-1126-48c4-8efd-9c3687de8524/17-MARCOS_ALONSO.jpg?width=470&height=470"},
	{Name: "Kounde", Age: "26", Number: "23", Position: "Defender", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/66fdb7c1-455d-45bb-a132-060fca97925e/23-JULES_KOUNDE.jpg?width=470&height=470"},
	{Name: "Gavi", Age: "18", Number: "6", Position: "Midfielder", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/1a22465f-7e9b-463d-981a-908dd221dd88/30-GAVI_.jpg?width=470&height=470"},
	{Name: "Pedri", Age: "19", Number: "8", Position: "Midfielder", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/4b341216-a50a-4d80-95ff-5ca2e5dcf640/08-PEDRI.jpg?width=470&height=470"},
	{Name: "Lopez", Age: "17", Number: "16", Position: "Midfielder", Image: "https://www.fcbarcelona.com/photo-resources/2024/02/01/4594c085-fa9b-4900-816a-8e7c99e2e625/mini_16-FERMIN.jpg?width=470&height=470"},
	{Name: "Romeu", Age: "30", Number: "18", Position: "Midfielder", Image: "https://www.fcbarcelona.com/photo-resources/2024/02/01/4594c085-fa9b-4900-816a-8e7c99e2e625/mini_16-FERMIN.jpg?width=470&height=470"},
	{Name: "Roberto", Age: "34", Number: "20", Position: "Midfielder", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/b30938f8-9ce1-4a72-be4a-bac6346abf13/20-S_ROBERTO.jpg?width=470&height=470"},
	{Name: "De Jong", Age: "28", Number: "21", Position: "Midfielder", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/d2214796-ebc8-4b1c-bb5d-37f2ce5883c7/21-FRENKIE_DE_JONG.jpg?width=470&height=470"},
	{Name: "Gundogan", Age: "32", Number: "22", Position: "Midfielder", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/3f3b9374-b75e-43f7-9e3e-ceddd9dc211d/jugador_fitxa-gundogan.jpg?width=470&height=470"},
	{Name: "Torres", Age: "25", Number: "7", Position: "Forward", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/166959ba-fb56-46ae-842b-e778d072ac5a/11-FERRAN_TORRES.jpg?width=470&height=470"},
	{Name: "Lewandowski", Age: "34", Number: "9", Position: "Forward", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/04101be1-2cef-4fb8-ae97-99a99297f65c/09-ROBERT_LEWANDOWSKI.jpg?width=470&height=470"},
	{Name: "Raphinha", Age: "24", Number: "11", Position: "Forward", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/865c65cd-9dc0-41b8-b373-6a922e29d594/22-RAPHINHA.jpg?width=470&height=470"},
	{Name: "Felix", Age: "23", Number: "14", Position: "Forward", Image: "https://www.fcbarcelona.com/photo-resources/2023/10/03/8a21c5f5-bea0-470e-bd2b-c905aeab7ade/jugador_fitxa-felix.jpg?width=470&height=470"},
	{Name: "Roque", Age: "19", Number: "19", Position: "Forward", Image: "https://www.fcbarcelona.com/photo-resources/2024/01/05/1f5ecf24-2a9e-4b96-af12-d22412ee58d2/19-VITOR_ROQUE.jpg?width=470&height=470"},
}

func GetAllPlayers() []Player {
	return players
}

func GetPlayerByName(name string) Player {
	for i := 0; i < len(players); i++ {
		player := players[i]
		nameWithoutSpaces := strings.ReplaceAll(player.Name, " ", "")
		nameWithoutSpaces = strings.ReplaceAll(nameWithoutSpaces, "/", "")
		if nameWithoutSpaces == name {
			return player
		}
	}
	return Player{}
}
