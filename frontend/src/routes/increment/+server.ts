import { json } from '@sveltejs/kit';

export async function POST() {
    try {
        console.log("Sending request to backend...");

        const response = await fetch('http://localhost:8080/counter.v1.CounterService/Increment', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({})
        });

        console.log("Response status:", response.status);

        if (!response.ok) {
            throw new Error(`Backend request failed with status ${response.status}`);
        }

        const data = await response.json();
        console.log("Backend responded:", data);

        if (!("newValue" in data)) {
            console.error("Response is missing 'newValue' field!", data);
        }

        return json({ newValue: data.newValue }); 
    } catch (error) {
        console.error("Error calling backend:", error);
        return json({ error: 'Failed to increment counter' }, { status: 500 });
    }
}
