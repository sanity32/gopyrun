import os


print("{{.Greetings}}")

open(f"{{.Filename}}", mode='a').close()

print(os.listdir(os.curdir) )