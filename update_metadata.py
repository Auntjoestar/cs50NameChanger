import json
import sys

# Define your metadata replacements here
metadata = {
    "{{.Info.ProductVersion}}": "1.0.0",
    "{{.Info.CompanyName}}": "Auntjoestar",
    "{{.Info.ProductName}}": "CS50 Name Changer",
    "{{.Info.Copyright}}": "© 2024 Auntjoestar. All Rights Reserved.",
    "{{.Info.FileDescription}}": "Cambiador de nombre para imagenes destinadas al SmugMug de CS50 Nicaragua.",
    "{{.Info.Comments}}": "Cambiador de nombre para imagenes destinadas al SmugMug de CS50 Nicaragua."
}

def update_json_file(file_path, replacements):
    try:
        # Read the JSON file
        with open(file_path, "r", encoding="utf-8") as f:
            data = json.load(f)

        # Convert JSON object to a string for replacements
        json_string = json.dumps(data)

        # Replace placeholders
        for placeholder, value in replacements.items():
            json_string = json_string.replace(placeholder, value)

        # Parse back to JSON object
        updated_data = json.loads(json_string)

        # Write updated JSON back to the file
        with open(file_path, "w", encoding="utf-8") as f:
            json.dump(updated_data, f, indent=4)

        print(f"Updated '{file_path}' successfully!")

    except Exception as e:
        print(f"Error updating JSON file: {e}")
        sys.exit(1)

# Update the file
if __name__ == "__main__":
    # Path to your JSON file (e.g., wails.json)
    json_file_path = "wails.json"

    update_json_file(json_file_path, metadata)
