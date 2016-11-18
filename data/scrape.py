"""Download and scrape primes from UTM website"""

from io import BytesIO
from zipfile import ZipFile
import requests
import sqlite3

conn = sqlite3.connect('primes.db')
db = conn.cursor()

def init_db():
    """Initialize to check create"""
    # Create table
    try:
        db.execute('''CREATE TABLE numbers
                    (id integer primary key, num text)''')
    except Exception:
        # Table and Database has already been initalized
        pass

BASE_URL = "http://primes.utm.edu/lists/small/millions/primes{}.zip"

def unzip_file(file_url):
    """Downloads a file over the network and unzips it"""
    url = requests.get(file_url)
    save_file(url)
    zipfile = ZipFile(BytesIO(url.content))
    zip_names = zipfile.namelist()
    # if unzipped return 1
    if len(zip_names) == 1:
        file_name = zip_names.pop()
        extracted_file = zipfile.open(file_name)
        return extracted_file

def save_file(req):
    """Saves a file from file context"""
    local_filename = req.url.split('/')[-1]
    with open(local_filename, 'wb') as f:
        for chunk in req.iter_content(chunk_size=1024):
            if chunk: # filter out keep-alive new chunks
                f.write(chunk)
                #f.flush() commented by recommendation from J.F.Sebastian
            return local_filename

def download_files(start=1, end=3):
    """Iterate over the primes on the site"""
    for page_num in range(start, end):
        parse_zip_file(unzip_file(BASE_URL.format(page_num)))
    pass

def parse_zip_file(inputfile):
    """Reads and converts byte buffet in file input"""
    for line in inputfile.readlines():
        if "primes.utm.edu" in str(line):
            continue
        for prime in line.decode('utf-8').split():
            db.execute("INSERT INTO numbers(num) VALUES({})".format(str(prime)))
    pass


def main():
    init_db()
    download_files()
    pass

if __name__ == '__main__':
    import sys
    sys.exit(int(main() or 0))

