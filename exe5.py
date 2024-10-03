def inverte(string: str):
    string_reversa: list[str] = [string[len(string) - 1 - i] for i in range(len(string))]

    print("".join(string_reversa))

inverte("pato caiu na lagoa")
