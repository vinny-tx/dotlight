import { useState } from 'react';

interface CreateGoalProps {
  setGoals: React.Dispatch<React.SetStateAction<any[]>>;
}

const CreateGoal = ({ setGoals }: CreateGoalProps) => {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [dueDate, setDueDate] = useState('');
  const [status, setStatus] = useState('Not Started');
  
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!name || !description || !dueDate) {
      alert("All fields are required!");
      return;
    }

    const newGoal = {
      title: name,
      description,
      dueDate,
      status,
    };

    const response = await fetch("http://localhost:8080/api/goals/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newGoal),
    });

    if (response.ok) {
      const updatedGoals = await fetch('http://localhost:8080/api/goals');
      const updatedGoalsData = await updatedGoals.json();
      setGoals(updatedGoalsData);  // Update the goals state in App.tsx
      alert("Goal created successfully!");
      setName('');
      setDescription('');
      setDueDate('');
      setStatus('Not Started');
    } else {
      alert("Error creating goal");
    }
  };

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">Create a New Goal</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <input
          type="text"
          placeholder="Goal Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="w-full p-2 border border-gray-300 rounded-md"
        />
        <textarea
          placeholder="Goal Description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className="w-full p-2 border border-gray-300 rounded-md"
        />
        <input
          type="date"
          value={dueDate} // This input will have a value in "YYYY-MM-DD" format
          onChange={(e) => setDueDate(e.target.value)}
          className="w-full p-2 border border-gray-300 rounded-md"
        />
        <select
          value={status}
          onChange={(e) => setStatus(e.target.value)}
          className="w-full p-2 border border-gray-300 rounded-md"
        >
          <option>Not Started</option>
          <option>In Progress</option>
          <option>Completed</option>
        </select>
        <button
          type="submit"
          className="w-full bg-blue-500 py-2 rounded-md hover:bg-blue-600"
        >
          Create Goal
        </button>
      </form>

    </div>
  );
};

export default CreateGoal;
