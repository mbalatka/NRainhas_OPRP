import matplotlib.pyplot as plt

def importaDados(arquivo):
    f = open(arquivo)
    linhas = f.readlines()
    valores = []
    for line in linhas:
        valores.append(float(line))
    print(valores)
    f.close()
    return sum(valores)/float(len(valores))

def geraGraficos(dados):
    fig, ax = plt.subplots(1, 1)
    plt.xlim(1, 6)
    plt.ylim(1, 6)
    ax.set_xlabel("Quantidade de Goroutines")
    ax.set_ylabel("SpeedUp")
    ax.set_title("N-Rainhas 6x6 força bruta")
    ax.plot(dados['th'], dados['tempo'], 'o-')
    plt.show()


d = {}
d['th'] = [1, 2, 3, 6]
d['tempo'] = []

d['tempo'].append(importaDados('size6_1goroutine.csv')/importaDados('size6_1goroutine.csv'))
d['tempo'].append(importaDados('size6_1goroutine.csv')/importaDados('size6_2goroutine.csv'))
d['tempo'].append(importaDados('size6_1goroutine.csv')/importaDados('size6_3goroutine.csv'))
d['tempo'].append(importaDados('size6_1goroutine.csv')/importaDados('size6_6goroutine.csv'))


geraGraficos(d)
