# ping_test.py
import subprocess

def perform_ping(address):
    try:
        subprocess.run(["ping", "-c", "4", address], check=True)
        print("Ping successful to", address)
    except subprocess.CalledProcessError:
        print("Ping failed to", address)

perform_ping("google.com")