const express = require('express');
const fs = require('fs/promises');

const app = express();
const port = 3000;

const dataFilePath = 'data.json';

// Middleware to parse JSON in the request body
app.use(express.json());

app.get('/', (req, res) => {
    res.send('Welcome to the Node.js CRUD server with JSON storage!');
});

// Read data from the JSON file
const readData = async () => {
  try {
    const data = await fs.readFile(dataFilePath, 'utf8');
    return JSON.parse(data);
  } catch (error) {
    return [];
  }
};

// Write data to the JSON file
const writeData = async (data) => {
  await fs.writeFile(dataFilePath, JSON.stringify(data, null, 2));
};

// Get all items

app.get('/items', async (req, res) => {
  const data = await readData();
  res.json(data);
});

// Get a specific item by ID
app.get('/items/:id', async (req, res) => {
  const data = await readData();
  const itemId = parseInt(req.params.id, 10);
  const item = data.find((item) => item.id === itemId);

  if (item) {
    res.json(item);
  } else {
    res.status(404).json({ error: 'Item not found' });
  }
});

// Create a new item
app.post('/items', async (req, res) => {
  const data = await readData();
  const newItem = req.body;
  newItem.id = data.length + 1;
  data.push(newItem);

  await writeData(data);

  res.status(201).json(newItem);
});

// Update an item by ID
app.put('/items/:id', async (req, res) => {
  const data = await readData();
  const itemId = parseInt(req.params.id, 10);
  const updatedItemIndex = data.findIndex((item) => item.id === itemId);

  if (updatedItemIndex !== -1) {
    data[updatedItemIndex] = { ...data[updatedItemIndex], ...req.body };

    await writeData(data);

    res.json(data[updatedItemIndex]);
  } else {
    res.status(404).json({ error: 'Item not found' });
  }
});

// Delete an item by ID
app.delete('/items/:id', async (req, res) => {
  const data = await readData();
  const itemId = parseInt(req.params.id, 10);
  const remainingItems = data.filter((item) => item.id !== itemId);

  if (remainingItems.length < data.length) {
    await writeData(remainingItems);

    res.json({ message: 'Item deleted successfully' });
  } else {
    res.status(404).json({ error: 'Item not found' });
  }
});

// Start the server
app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});
