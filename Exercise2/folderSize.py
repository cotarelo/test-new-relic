# folderSize.py

# Displays the size of a directory passed as a parameter
# The parameter is passed as first argument in text
# Example python folderSize.py /peth/to/directory

import os
import sys


def getFolderSize(folder):
    total_size = os.path.getsize(folder)
    for item in os.listdir(folder):
        itempath = os.path.join(folder, item)
        if os.path.isfile(itempath):
            total_size += os.path.getsize(itempath)
        elif os.path.isdir(itempath):
            total_size += getFolderSize(itempath)
    return total_size

print ("Size: " + str(getFolderSize(sys.argv[1])))
