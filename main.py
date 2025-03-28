def text_search (text):
    count = 0
    for el in text:
        if el == "." or el == ",":
            count += 1
    return count

def main():
    # input didn't work for me, so that's it :)
    text = "On Thursday the fourth at four and a quarter o'clock the Ligurian traffic controller was regulating in Liguria, but thirty-three ships tacked and tacked, but still did not catch."
    count = text_search(str(text))
    print(count)

main()