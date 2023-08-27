import { useState } from 'react';
import UpdatedProvider from '@/components/UpdatedProvider';
import { FaPencilAlt } from "react-icons/fa";

export default function Home() {
  const initialData = {
    id: '123',
    name: 'Product Name',
    webhook: 'https://webhook.example.com',
    provider: 'Provider Name',
  };

  const [showUpdate, setShowUpdate] = useState(false);
  const [selectedData, setSelectedData] = useState(initialData);

  const handleEdit = (data) => {
    setSelectedData(data);
    setShowUpdate(true);
  };

  const handleUpdate = (updatedData) => {
    console.log('Updated data:', updatedData);
    // Perform update logic here (e.g., send API request)
    // After update, fetch new data and update the UI
    setSelectedData(updatedData);
    setShowUpdate(false);
  };

  const handleCancel = () => {
    setShowUpdate(false);
  };

  return (
    <div>
      {/* Render your data list as a table */}
      {/* <table className="border-collapse w-full"> */}
        {/* <thead>
          <tr>
            <th className="border p-2 text-black">ID</th>
            <th className="border p-2 text-black">Name</th>
            <th className="border p-2 text-black">Webhook</th>
            <th className="border p-2 text-black">Provider</th>
            <th className="border p-2 text-black">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td className="border p-2">{selectedData.id}</td>
            <td className="border p-2">{selectedData.name}</td>
            <td className="border p-2">{selectedData.webhook}</td>
            <td className="border p-2">{selectedData.provider}</td>
            <td className="border p-2">
              <button onClick={() => handleEdit(selectedData)}>
                <FaPencilAlt className="text-blue-500" />
              </button>
            </td>
          </tr> */}
          {/* Render other data items */}
        {/* </tbody>
      </table> */}
<button onClick={() => handleEdit(selectedData)}>
                <FaPencilAlt className="text-blue-500" />
              </button>
            {/* </td> */}
      {/* Render the update form popup */}
      {showUpdate && (
        <UpdatedProvider
          data={selectedData}
          onUpdate={handleUpdate}
          onCancel={handleCancel}
        />
      )}
    </div>
  );
}
