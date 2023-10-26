# Lorem ipsum generator

## Command line options

One of the below, if multiple specified, the highest order takes priority:
-p N generate n paragraphs,
-s N generate n sentences in one paragraph,
-w N generate n words in one (?) paragraph,
-c N generate n non-whitespace characters in one (?) paragraph,

where N is a positive integer

Additional options
-d          default beginning of the text, ex. Lorem ipsum...,
            taken from the first line of the text file
-l [file]   load language pack

## Output modes

1. Text: paragraphs ended/separated with '\n\n'.
2. Document with headings, lists, etc. (options?)
   a. Markdown
   b. HTML

## Generation method

### By paragraphs

If number of paragraphs `p` specified for more than default 1, generate `p`
paragraphs. The number of sentences in each paragraph is randomly chosen with a
normal distribution apprioximation. Then sentences are taken from the resource
file and shuffled with no repeating within a given distance (?), ex. 1/4 of the
collection size.

### By sentences (-> one paragraph)

The procedure of choosing sentences is the same for specified `s` number, but
only one paragraph is created.

### By words (-> one paragraph)

If `w` is specified, sentences are added until the number of words is at least
`w`, and then words are removed if needed.

### By characters (-> one paragraph)

Character count option also takes into the account the length of the words and
replaces the words with no repetitions until the desired length is obtained
(might be exceeded the easy method wouldn't yield a desirable result).

