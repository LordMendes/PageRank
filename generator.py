import os
import random
import sys
import csv

def load_words_from_csv(filename):
    words = []
    with open(filename, newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            words.append(row['word'])
    return words

def generate_random_text(words, num_words):
    return ' '.join(random.choices(words, k=num_words))

def generate_random_tags(words, num_tags):
    return random.sample(words, num_tags)

def generate_html_file(n, links, words, folder):
    filename = os.path.join(folder, f"website{n}.html")
    random_text = generate_random_text(words, 50)
    random_tags = generate_random_tags(words, 5)
    
    with open(filename, "w") as f:
        f.write("<html>\n")
        f.write("<head>\n")
        f.write(f"<title>Website {n}</title>\n")
        for tag in random_tags:
            f.write(f"<meta name=\"{tag}\" content=\"{tag}\">\n")
        f.write("</head>\n")
        f.write("<body>\n")
        f.write(f"<h1>Website {n}</h1>\n")
        f.write(f"<p>{random_text}</p>\n")
        for link in links:
            f.write(f"<a href=\"{link}\">{link}</a><br>\n")
        f.write("</body>\n")
        f.write("</html>\n")

def generate_websites(N, words, folder):
    all_links = [f"website{i}.html" for i in range(1, N+1)]
    for i in range(1, N+1):
        links = all_links.copy()
        links.remove(f"website{i}.html")  # Remove self-link
        random.shuffle(links)
        num_links = random.randint(1, N-1)
        selected_links = random.sample(links, num_links)
        generate_html_file(i, selected_links, words, folder)

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python generator.py N words_pos.csv")
        sys.exit(1)
    
    N = int(sys.argv[1])
    if N < 2:
        print("N must be greater than 1")
        sys.exit(1)
    
    words_file = sys.argv[2]
    words = load_words_from_csv(words_file)
    
    # Create the websites folder if it does not exist
    output_folder = "websites"
    if not os.path.exists(output_folder):
        os.makedirs(output_folder)
    
    generate_websites(N, words, output_folder)
    print(f"Generated {N} websites in the '{output_folder}' folder")
    sys.exit(0)
