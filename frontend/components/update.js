import React, { useState, useEffect } from 'react'
import Head from 'next/head'
import { useRouter } from 'next/router'
import axios from 'axios'
import {FcCheckmark} from 'react-icons/fc'

export default function Update() {
    const router = useRouter();
    
    const [ID, setID] = useState('');
    const [name, setName] = useState('');
    const [link, setLink] = useState('');
    const[int, setInt] = useState('');

    const[showmodal, setShowModal] = useState(false);
    
    // useEffect(() =>{
    //     const fetchData = async ()=> {
    //         try {
    //             const
    //         }
    //     }
    // })

    return(
       <div>
        {showmodal && (
          <div className='fixed inset-0 flex items-center justify-center z-50 bg-black opacity-75'>
            <div className='bg-slate-700 p-8 rounded-lg flex flex-col text-white'>
              <h2 className='text-xl font-semibold mb-4 text-white text-center'>Product Successfully Updated!</h2>
              {/* <p className='mb-2'>Title : {updatedTitle}</p>
              <p className='mb-4'>Price : {updatedPrice}</p> */}
              <FcCheckmark className="text-6xl text-emerald-500 animate-pulse mx-auto mb-4"/>
              <button
                onClick={() => setShowModal(false)} 
                className='bg-emerald-600 px-4 py-2 hover:bg-emerald-400 font-bold rounded-lg mt-4 text-white'
              >
                Close
              </button> 
            </div>
          </div>
          )}
       </div>
    )

}