interface Goal {
  id: string;
  title: string;
  description: string;
  dueDate: string;
  status: string;
}

const GoalList = ({ goals }: { goals: Goal[] }) => {
  return (
    <div className="mt-8">
      <h2 className="text-xl font-semibold mb-4">Goals</h2>
      <ul>
        {goals.map((goal) => (
          <li key={goal.id} className="mb-2">
            <p className="font-bold">{goal.title}</p>
            <p>{goal.description}</p>
            <p>Due Date: {goal.dueDate}</p>
            <p>Status: {goal.status}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default GoalList;
