import { useState, useEffect } from "react";
import { askDotlight } from "./api";
import GoalsList from "./components/Goals/GoalsList"; 
import CreateGoal from "./components/Goals/CreateGoal";

export default function App() {
  const [question, setQuestion] = useState("");
  const [answer, setAnswer] = useState("");
  const [chatHistory, setChatHistory] = useState<{ question: string; answer: string }[]>([]);
  const [goals, setGoals] = useState<any[]>([]);

  useEffect(() => {
    const fetchGoals = async () => {
      const response = await fetch('http://localhost:8080/api/goals');
      const data = await response.json();
      setGoals(data);
    };

    fetchGoals();
  }, []);

  const handleAsk = async () => {
    try {
      const result = await askDotlight(question);
      setAnswer(result);
      setChatHistory([...chatHistory, { question, answer: result }]);
    } catch (err) {
      setAnswer("Error: " + (err as Error).message);
    }
  };

  return (
    <div className="min-h-screen bg-gray-900 p-8 text-white">
      <h1 className="text-3xl font-bold mb-6">Dotlight</h1>
      
      {/* Chat History Section */}
      <div className="bg-gray-800 p-4 rounded-md mb-6">
        <h2 className="text-xl font-semibold mb-4">Chat History</h2>
        <div className="space-y-4">
          {chatHistory.map((chat, index) => (
            <div key={index} className="space-y-2">
              <div className="font-semibold">Q: {chat.question}</div>
              <div>A: {chat.answer}</div>
            </div>
          ))}
        </div>
      </div>
      
      {/* Chat Input Section */}
      <div className="mb-6">
        <input
          value={question}
          onChange={(e) => setQuestion(e.target.value)}
          className="w-full p-2 rounded-md mb-2"
          placeholder="Ask me anything..."
        />
        <button
          onClick={handleAsk}
          className="w-full bg-blue-500 py-2 rounded-md hover:bg-blue-600"
        >
          Ask
        </button>
        <div className="mt-4 bg-gray-800 p-4 rounded-md">{answer}</div>
      </div>
      
       <CreateGoal setGoals={setGoals} />
      
      <GoalsList goals={goals} />
    </div>
  );
}
