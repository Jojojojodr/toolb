import hashlib, sys

passwords = []

args = " ".join(sys.argv[1:])
args = args.split(" ")
hash_method = args[0]
hash = args[1]

print(f"Hash method: {hash_method}")

with open("passwords.txt", "r") as f:
    passwords = f.read().splitlines()
    
for password in passwords:
    match hash_method:
        case "md5":
            hashed_password = hashlib.md5(password.encode('utf-8')).hexdigest()
            if hashed_password == hash:
                print(f"Password found: {password}")
                break
        case "sha1":
            hashed_password = hashlib.sha1(password.encode('utf-8')).hexdigest()
            if hashed_password == hash:
                print(f"Password found: {password}")
                break
        case "sha256":
            hashed_password = hashlib.sha256(password.encode('utf-8')).hexdigest()
            if hashed_password == hash:
                print(f"Password found: {password}")
                break
        case _:
            print("Invalid hash method")
            break