def serch_in_text(text):
    x = 0
    for el in text:
        if el == "." or el == ",":
            x += 1
    return x

def main():
    text = "wdl,aw,da[,d,[da,wd,a]]"
    x = serch_in_text(text)
    print(x)
    print(__builtins__)

main()