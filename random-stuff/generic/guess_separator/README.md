# Guessing the separator in files

In certain cases we might receive a whole bunch of data files for processing where each file has one record per line, but where different files might use different separators (e.g., `tabs` or `whitespace` or `“*”s`).

To be able to process such files in bulk we need to be able to determine the *separator* used for each file. The *guess_separator* attempts to identify the separator for the file it is given to work on.
