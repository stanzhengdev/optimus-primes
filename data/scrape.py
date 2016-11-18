"""Download and scrape primes from UTM website"""

from io import BytesIO
from zipfile import ZipFile
import requests

BASE_URL = "http://primes.utm.edu/lists/small/millions/primes{}.zip"

def unzip_file(file_url):
    """Downloads a file over the network and unzips it"""
    url = requests.get(file_url)
    zipfile = ZipFile(BytesIO(url.content))
    zip_names = zipfile.namelist()
    # if unzipped return 1
    if len(zip_names) == 1:
        file_name = zip_names.pop()
        extracted_file = zipfile.open(file_name)
        return extracted_file

def download_files(start=0, end=50):
    """Iterate over the primes on the site"""
    pass

def main():
    pass

if __name__ == '__main__':
    import sys
    sys.exit(int(main() or 0))

