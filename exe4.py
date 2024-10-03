faturas: list[dict[str, str | float]] = [
    {
        "estado": "SP",
        "faturamento": 67836.43
    },
    {
        "estado": "RJ",
        "faturamento": 36678.66
    },
    {
        "estado": "MG",
        "faturamento": 29229.88
    },
    {
        "estado": "ES",
        "faturamento": 27165.48
    },
    {
        "estado": "Outros",
        "faturamento": 19849.53
    }
]
total: float = 0.0

for fatura in faturas:
    assert fatura.get("faturamento") is not None
    total += fatura.get("faturamento")

for fatura in faturas:
    assert fatura.get("estado") is not None and fatura.get("faturamento") is not None
    print(f"O estado de {fatura.get("estado")} teve uma partipação de {(fatura.get("faturamento") / total) * 100}%")
