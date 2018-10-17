#!/usr/bin/env python3
"""Script to collapse multiple .go files into one.

Handles package statements but not imports.
"""

import os
import subprocess

from progress.bar import ShadyBar


def derive_new_filename(filename):
    """Derives a non-numbered filename from a possibly numbered one."""
    while filename[-1].isdigit():
        filename = filename[:-1]
    return filename


def main():
    go_files = [filename[:-3] for filename in os.listdir('iso20022') if filename.endswith('.go')]
    go_files.sort()
    for filename in ShadyBar('Munging files').iter(go_files):
        new_filename = os.path.join('iso20022', derive_new_filename(filename) + '.go')
        filename = os.path.join('iso20022', filename + '.go')
        if new_filename == filename:
            continue
        if not os.path.exists(new_filename):
            os.rename(filename, new_filename)
        else:
            with open(filename) as oldf:
                with open(new_filename, 'a') as newf:
                    for line in oldf:
                        if not line.startswith('package'):
                            newf.write(line)
            os.unlink(filename)


if __name__ == '__main__':
    main()
