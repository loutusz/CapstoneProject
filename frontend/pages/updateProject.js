import { useState,useEffect } from 'react';
import UpdatedProject from '@/components/UpdatedProject';
import { useRouter } from 'next/router';
import { FaPencilAlt } from "react-icons/fa";
import axios from 'axios';

export default function Home() {
  const router = useRouter()
  const project_id = router.query.project_id
  const [data, setData] = useState([])

//   const initialData = {
//     id: '123',
//     name: 'Product Name',
//     webhook: 'https://webhook.example.com',
//     provider: 'Provider Name',
//   };

  const [showUpdate, setShowUpdate] = useState(false);
//   const [selectedData, setSelectedData] = useState(data);

  useEffect(() => {
    axios.get(`http://localhost:8050/project/id/${project_id}`)
    .then(res => setData(res.data))
    .catch(err => console.log(err));
    // const fetchDataProduct = async () => {
    //   try {
    //     const resp = await axios.get(`https://dummyjson.com/products/${id}`);

    //   } catch (error) {
    //     console.log('Error Fetch Data', error)
    //   }
    // }
    // fetchDataProduct();
  },[project_id]);

  const handleEdit = (data) => {
    // setSelectedData(data);
    setData(data);
    setShowUpdate(true);
  };

  const handleSubmit = async (e) => {  
    try {
        const updatedData = {
            id: data.id,
            projectname: data.name,
            webhook: data.webhook,
            provider: data.provider,
        };

        await axios.put(`http://localhost:8050/project/edit/{{project_id}}`, updatedData);

        setData(updatedData);
    } catch (err) {
        console.log('Update error',err)
    }

    // console.log('Updated data:', updatedData);

    setShowUpdate(false);
  };

  const handleCancel = () => {
    setShowUpdate(false);
  };

  return (
    <div>
    <button onClick={() => handleEdit(data)}>
                    <FaPencilAlt className="text-blue-500" />
                </button>
                {/* </td> */}
        {/* Render the update form popup */}
        {showUpdate && (
            <UpdatedProject
            data={data }
            onUpdate={handleSubmit}
            onCancel={handleCancel}
            />
        )}
        </div>
    );
}