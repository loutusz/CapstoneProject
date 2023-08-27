import React, { useState } from 'react';

export default function UpdatedProvider({ data, onUpdate, onCancel }) {
  const [id, setId] = useState(data.id);
  const [projectname, setProjectName] = useState(data.name);
  const [webhook, setWebhook] = useState(data.webhook);
  const [provider, setProvider] = useState(data.provider);

  const handleUpdate = () => {
    const updatedData = {
      id,
      projectname,
      webhook,
      provider,
    };
    onUpdate(updatedData);
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center z-50 bg-black bg-opacity-70" >
    <div className="bg-white rounded-lg absolute max-w-[30rem] p-3 top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
      <div className="flex justify-between items-center mb-3">
        <h1 className="text-lg text-black mb-2">Update Provider</h1>
        <button className="text-black hover:text-red-600 mb-2" type="button" onClick={onCancel}>
          Close
        </button>
      </div>
      <table className="w-full mb-5 ">
        <tbody className=''>

          {/* ID*/}
          <tr>
            <td className="text-black font-semibold ">ID</td>
            <td className='pl-3'>
              <input
                type="text"
                value={id}
                disabled
                onChange={(e) => setId(e.target.value)}
                className="w-full border p-2 rounded outline-none"
              />
            </td>
          </tr>
          
        {/* Project Name */}
          <tr>
            <td className="text-black font-semibold">Project Name</td>
            <td className='pl-3'>
              <input
                type="text"
                value={projectname}
                onChange={(e) => setProjectName(e.target.value)}
                className="w-full border p-2 rounded outline-none"
              />
            </td>
          </tr>

          {/* Link Webhook */}
          <tr>
            <td className="text-black font-semibold"> Link Webhook</td>
            <td className ="pl-3"> 
              <input
                type="text"
                value={webhook}
                onChange={(e) => setWebhook(e.target.value)}
                className="w-full border p-2 rounded outline-none"
              />
            </td>
          </tr>

          {/* Nama Provider */}
          <tr>
            <td className="text-black font-semibold">Provider</td>
            <td className='pl-3'>
              <input
                type="text"
                value={provider}
                onChange={(e) => setProvider(e.target.value)}
                className="w-full border p-2 rounded outline-none"
              />
            </td>
          </tr>
        </tbody>
      </table>

      {/* Update */}
      <div className="flex justify-center items-center w-full space-x-3">
        <button
          type="button"
          className="px-4 py-1 bg-emerald-400 text-emerald-900 hover:bg-emerald-500 hover:text-white rounded-lg"
          onClick={handleUpdate}
        >
          Save
        </button>

      {/* Cancel */}
        <button
          type="button"
          className="px-4 py-1 bg-red-400 text-red-900 hover:bg-red-500 hover:text-white rounded-lg"
          onClick={onCancel}
        >
          Cancel
        </button>
      </div>
    </div>
    </div>
  );
}