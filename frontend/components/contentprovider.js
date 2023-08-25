import React, { useEffect, useState } from "react";
import axios from "@/pages/api/axios";

export default function ContentProvider() {
    const [provider, setProvider] = useState([]);
    const [isLoading, setIsLoading] = useState(true);

    const fetchData = async () => {
        try {
            const response = await axios.get('http://localhost:8050/project/all');
            setProvider(response.data.provider);
            setIsLoading(false);
        } catch (error) {
            console.error('Error fetching data: ', error);
        }
    };

    useEffect(() => {
        fetchData();
    }, []);

    return (
        <div className="flex justify-center mt-10 h-screen">
            {isLoading ? (
                <p>Loading...</p>
            ) : (
                <table>
                    <thead>
                        <tr>
                            {/* <th>
                                <label>
                                    <input type="checkbox" className="checkbox"/>
                                </label>
                            </th> */}
                            <th className="border-collapse border w-4/5 border-gray-800">ID</th>
                            <th className="border-collapse border w-4/5 border-gray-800">Project Name</th>
                            <th className="border-collapse border w-4/5 border-gray-800">Webhook Link</th>
                            <th className="border-collapse border w-4/5 border-gray-800">Integration</th>
                            <th className="border-collapse border w-4/5 border-gray-800">Edit</th>
                            <th className="border-collapse border w-4/5 border-gray-800">Delete</th>
                        </tr>
                    </thead>
                    <tbody>
                        {provider.map((item) => (
                             <tr key={item.id}>
                             <td className="px-6 py-4 whitespace-nowrap">{item.id}</td>
                             <td className="px-6 py-4 whitespace-nowrap">{item.name}</td>
                             <td className="px-6 py-4 whitespace-nowrap">{item.link}</td>
                             <td className="px-6 py-4 whitespace-nowrap">{item.provider}</td>
                             <td className="px-6 py-4 whitespace-nowrap">
                                {/* button edit*/}
                             </td>
                             <td className="px-6 py-4 whitespace-nowrap">
                                {/* button delete*/}
                             </td>
                             </tr>
                        ))}
                    </tbody>
                </table>
            )}
        </div>
    )
}