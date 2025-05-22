#This program converts a txt version on the bible into a more readable format

from dict import bible     #All dictionaries declared here


with open("test.txt") as file:
    # read through the input file, removing trailing newlines
    lines = [l.strip() for l in file.readlines()]

    counter = 0

    for line in lines:
        if not line:    #Avoids line breaks
            pass

        elif not str.isdigit(line[0]):      #Looks for a title
            book = line
        
        else:
            pass #IDK how to do this


print(bible)