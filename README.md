# Notes
- If I had more time I would like to save the assets into another folder.
- Make sure to include protocol in the URL (ex: https://www.google.com or https://github.com/nurcholisnanda)

# How to use
- BUILD
    - docker build . --tag fetch
- FETCH
    - docker run --volume $(pwd):/app fetch https://google.com https://github.com/nurcholisnanda
- PRINT METADATA
    - docker run --volume $(pwd):/app fetch --metadata https://google.com https://github.com/nurcholisnanda