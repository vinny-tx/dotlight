export async function askDotlight(question: string) {
    const res = await fetch("http://localhost:8080/api/ask", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ question })
    });
  
    if (!res.ok) {
      throw new Error(`API error: ${res.status}`);
    }
  
    const data = await res.json();
    return data.answer;
  }
  