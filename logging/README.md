# MdTF Image Metadata JSON Log File Format For Face Acquisition Systems

This repository contains the specification for a JSON log file format used to store metadata associated with a set of images. 
The metadata includes the ISO datetime when the images were captured and the station ID from which the images were captured.

## Metadata Format

The JSON log file contains the following structure:

```json
{
  "metadata_version": "1.0",
  "images": [
    {
      "filename": "string",
      "capture_datetime": "string (ISO 8601 format)",
      "station_id": "string"
    },
    ...
  ]
}
```

### Fields

- \`metadata_version\`: A string indicating the version of the metadata file format.
- \`images\`: An array of objects, each representing metadata for a single image.
  - \`filename\`: The filename of the image.
  - \`image_type\`: The type of image set to "Face".
  - \`capture_datetime\`: The ISO 8601 formatted datetime string indicating when the image was captured.
  - \`station_id\`: The identifier of the station from which the image was captured.

## Example

Here is an example of the JSON log file:

```json
{
  "metadata_version": "1.0",
  "images": [
    {
      "filename": "image_001.jpg",
      "image_type": "Face",
      "capture_datetime": "2024-06-21T10:15:30Z",
      "station_id": "Station_A"
    },
    {
      "filename": "image_002.jpg",
      "image_type": "Face",
      "capture_datetime": "2024-06-21T10:20:45Z",
      "station_id": "Station_A"
    },
    {
      "filename": "image_003.jpg",
      "image_type": "Face",
      "capture_datetime": "2024-06-21T10:25:00Z",
      "station_id": "Station_A"
    }
  ]
}
```

## Image Storage

All image files referenced in the JSON log file should be stored together in a single folder. 

### Example Directory Structure

The following is an example of how to structure your log and image files:

```
/project-directory
│
├── images/
│   ├── image_001.jpg
│   ├── image_002.jpg
│   └── image_003.jpg
│
└── metadata.json
```

- \`/project-directory\`: The root directory of your project.
- \`images/\`: A folder containing all the image files referenced in the \`metadata.json\` file.
- \`metadata.json\`: The JSON log file containing metadata for each image.

## Usage

1. **Create the JSON Metadata File:**
   - Follow the specified format to create your \`metadata.json\` file.
   - Populate the file with metadata for each image, including the filename, capture datetime in ISO 8601 format, and station ID.

2. **Store Images Together:**
   - Store all image files referenced in the \`metadata.json\` file in a single folder (e.g., \`images/\`).  Each image must have a unique filename.

3. **Ensure Accuracy:**
   - Verify that the \`filename\` field in the \`metadata.json\` file accurately points to the corresponding image files in the folder.
