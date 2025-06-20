const { createClient } = require("@supabase/supabase-js");

const supabaseUrl = process.env.SUPABASE_URL;
const supabaseKey = process.env.SUPABASE_KEY;
const supabase = createClient(supabaseUrl, supabaseKey);

export default async function handler(req, res) {
  const { task } = req.body;

  if (!task) return res.status(400).send("Missing task");

  const { data, error } = await supabase.from("todos").insert([{ task }]);

  if (error) return res.status(500).send(error.message);

  res.send(`Added task: ${task}`);
}
