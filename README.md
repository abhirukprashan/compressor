Go-RLE: A Learning Experiment in Compression
⚠️ DISCLAIMER: DO NOT USE FOR PRODUCTION

This is an educational project created to study the mechanics of Run-Length Encoding (RLE). This tool is not optimized for performance or safety. Because of the nature of RLE, certain file types (like images or already compressed files) will actually increase in size when processed by this tool.
🧐 What is this?

This project is a simple command-line utility written in Go that implements a basic Run-Length Encoding algorithm.

In a "Best Case" scenario (like a text file with 1,000 "A" characters), this tool can shrink the file significantly. In a "Worst Case" scenario (like a photo or random data), the file size will double because the algorithm adds a "count" byte to every single unique character.
How it Works (The Logic)

The algorithm iterates through the byte stream and looks for consecutive identical bytes:

    It counts how many times a byte repeats (up to 255).

    It writes the Count followed by the Value.

Example:

    Input: AAAAABBB ( bytes)

    Output: [5]A[3]B (4 bytes) — 50% Compression

The "Size Increase" Trap:

    Input: ABC (3 bytes)

    Output: [1]A[1]B[1]C (6 bytes) — 200% Increase

🛠 Features

    Compression: Converts repeating sequences into [count][value] pairs.

    Decompression: Reconstructs the original data from the pairs.

    Safety Checks: Includes basic boundary checks to prevent crashes on malformed data.
