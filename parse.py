# Read
input = open("fortune.txt")
DELIMITER = '%\n'
fortunes = []
fortune = ""
for line in input:
    if line != DELIMITER:
        line = line.replace("\"", "\\\"")
        line = line.replace("'", "\'")
        line = line.replace("\n", "\\n")
        if "`" in line: # TODO weird incomplete lines because of this, create an exclusive if for it
            continue
        fortune += line
    else:
        fortunes.append(fortune)
        fortune = ""
# Write
output = open("fortunes.go", "w")
output.write("package main\n\n")
output.write("var fortunes = []string{\n")
for fortune in fortunes:
    output.write("\"" + fortune + "\",\n")
output.write("}\n")
input.close()
