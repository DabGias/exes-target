dias: list[dict[str, float]] = [
	{
		"dia": 1,
		"valor": 22174.1664
	},
	{
		"dia": 2,
		"valor": 24537.6698
	},
	{
		"dia": 3,
		"valor": 26139.6134
	},
	{
		"dia": 4,
		"valor": 0.0
	},
	{
		"dia": 5,
		"valor": 0.0
	},
	{
		"dia": 6,
		"valor": 26742.6612
	},
	{
		"dia": 7,
		"valor": 0.0
	},
	{
		"dia": 8,
		"valor": 42889.2258
	},
	{
		"dia": 9,
		"valor": 46251.174
	},
	{
		"dia": 10,
		"valor": 11191.4722
	},
	{
		"dia": 11,
		"valor": 0.0
	},
	{
		"dia": 12,
		"valor": 0.0
	},
	{
		"dia": 13,
		"valor": 3847.4823
	},
	{
		"dia": 14,
		"valor": 373.7838
	},
	{
		"dia": 15,
		"valor": 2659.7563
	},
	{
		"dia": 16,
		"valor": 48924.2448
	},
	{
		"dia": 17,
		"valor": 18419.2614
	},
	{
		"dia": 18,
		"valor": 0.0
	},
	{
		"dia": 19,
		"valor": 0.0
	},
	{
		"dia": 20,
		"valor": 35240.1826
	},
	{
		"dia": 21,
		"valor": 43829.1667
	},
	{
		"dia": 22,
		"valor": 18235.6852
	},
	{
		"dia": 23,
		"valor": 4355.0662
	},
	{
		"dia": 24,
		"valor": 13327.1025
	},
	{
		"dia": 25,
		"valor": 0.0
	},
	{
		"dia": 26,
		"valor": 0.0
	},
	{
		"dia": 27,
		"valor": 25681.8318
	},
	{
		"dia": 28,
		"valor": 1718.1221
	},
	{
		"dia": 29,
		"valor": 13220.495
	},
	{
		"dia": 30,
		"valor": 8414.61
	}
]
maior: float | None = dias[0].get("valor") if dias[0].get("valor") is not None else None
menor: float | None = dias[0].get("valor") if dias[0].get("valor") is not None else None
soma: float = 0.0
num_dias: int = 0

for dia in dias:
    assert dia.get("valor") is not None and maior is not None and menor is not None
    if dia.get("valor") != 0.0:
        soma += dia.get("valor")
        num_dias += 1

    if dia.get("valor") > maior:
        maior = dia.get("valor")
    elif dia.get("valor") < menor and dia.get("valor") != 0.0:
        menor = dia.get("valor")

media: float = soma / num_dias
dias_acima_da_media: int = 0

for dia in dias:
    assert dia.get("valor") is not None
    if dia.get("valor") > media:
        dias_acima_da_media += 1

print(f"Maior faturamento: {maior}")
print(f"Menor faturamento: {menor}")
print(f"Dias acima da média: {dias_acima_da_media}")
