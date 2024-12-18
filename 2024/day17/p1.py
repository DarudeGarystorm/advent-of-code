import os

script_dir = os.path.dirname(__file__)
file_path = os.path.join(script_dir, 'input.txt')

with open(file_path, "r") as infile:
    data = infile.read().splitlines()
    rA = int(data[0].split(': ')[1])
    rB = int(data[1].split(': ')[1])
    rC = int(data[2].split(': ')[1])
    instructions = [int(s) for s in data[4].split(': ')[1].split(',')]
    output = []
    
    def combo_operand(o):
        if operand == 4:
            return rA
        elif operand == 5:
            return rB
        elif operand == 6:
            return rC
        return o

    i = -2
    while i < len(instructions) - 2:
        i += 2
        opcode, operand = instructions[i], instructions[i+1]

        # opcode instructions
        if opcode == 0:
            rA = rA // 2**combo_operand(operand)
        elif opcode == 1:
            rB = rB ^ operand
        elif opcode == 2:
            rB = combo_operand(operand) % 8
        elif opcode == 3:
            if rA != 0:
                i = operand - 2
        elif opcode == 4:
            rB = rB ^ rC
        elif opcode == 5:
            output.append(combo_operand(operand) % 8)
        elif opcode == 6:
            rB = rA // 2**combo_operand(operand)
        elif opcode == 7:
            rC = rA // 2**combo_operand(operand)

    print(','.join(map(str, output)))
