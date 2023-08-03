import React, {useState} from 'react';
import Link from "next/link";
import Head from 'next/head';
import { FaRegEnvelope } from 'react-icons/fa';
import { MdPermIdentity, MdLockOutline } from 'react-icons/md';
import { BiSolidHide } from "react-icons/bi";

// export default function RegisterPage() {
  const useForm = () => {
    const [formData, setFormData] = useState({
        name: '',
        username: '',
        email: '',
        password: ''
      });
    
      const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
      };
    
      const handleSubmit = (e) => {
        e.preventDefault();
        // Lakukan aksi pendaftaran atau validasi form di sini
        // const fullname = /^[A-Z][a-zA-Z\s]+$/;
        // if(!fullname.test(formData.fullname)) {
        //   alert('Nama harus diawali dengan huruf kapital');
        //   return;
        // }

        const username = /^[A-Za-z]{5,15}$/;
        if(!username.test(formData.username)) {
          alert('Username harus terdiri dari 5-15 karakter');
          return;
        }
      
        const password = /^(?=.*[A-Z])(?=.*[0-9]).{6,10}$/
        if(!password.test(formData.password)) {
          alert('Password harus terdiri dari 6-10 karakter dengan minimal satu huruf kapital dan satu angka');
          return;
        }

        const email = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if(!email.test(formData.email)) {
          alert('Email harus memiliki format yang benar (contoh: example@example.com');
          return;
        }
        console.log('Pendaftaran berhasil');
        console.log(formData);
        // Reset form setelah pengiriman
        setFormData({
          name: '',
          username: '',
          email: '',
          password: ''
        });
      };

      return{
        formData, 
        handleChange, 
        handleSubmit,
      };
    };

    const SignUpPage = () => {
      const{formData, handleChange, handleSubmit} = useForm();
      const{show,showPass} = useState(false);

      return(
        <div>
            <div>
                <Head>
                <title>Sign Up</title>
                </Head>
            </div>

            {/* Backkground */}
            <div className='bg-gradient-to-r from-cyan-500/10 via-teal-300/10 to-sky-200/10 block h-screen items-center justify-center p-4 md:flex'>

               {/* Container */}
                <div className=' bg-white flex flex-col items-center max-w-screen-lg overflow-hidden rounded-lg space-y-8 shadow-[0_3px_10px_rgb(0,0,0,1)] w-full md:flex-row text-slate-700'>
                    
                    {/* Welcome*/}
                    <div className='bg-gradient-to-t from-sky-200 to-stone-50 flex flex-col items-center p-4 w-full md:w-2/5 text-white'>
                        <div className="flex flex-col justify-center items-start text-left text-slate-700 w-full p-10">
                         
                          <h2 className="text-3xl font-bold mb-4 ">Welcome To Jico</h2>
                          <p className='text-base font-normal mb-48'>
                            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam nec ultricies nisi.
                            Suspendisse pulvinar viverra nibh vel ultricies. Mauris tincidunt mollis diam, at mollis enim aliquet eget.
                            Fusce eros neque, pharetra eget tincidunt in, tincidunt nec tellus.
                          </p>

                          <div className='h-8 w-16 flex object-right-bottom'>
                            <img className='' src='https://upload.wikimedia.org/wikipedia/id/thumb/c/c4/Telkom_Indonesia_2013.svg/1200px-Telkom_Indonesia_2013.svg.png'/>
                            <img className='' src='https://cdn.icon-icons.com/icons2/2699/PNG/512/atlassian_jira_logo_icon_170512.png'/>
                          </div>
                        </div>       
                    </div>

                    {/* Sign Up */}
                        {/* Title */}
                        <div className='flex flex-col items-center justify-center w-full md:w-3/5 '>
                            <h1 className='text-slate-700 text-4xl font-semibold '>Sign Up</h1>
                        </div>

                        {/* Form */}
                        <form onSubmit={handleSubmit} className='flex flex-col items-center p-4 space-y-4'>
                            <div className='relative py-2'>

                                {/* Fullname */}
                                
                                {/* <div className="w-[80%] mx-auto bg-white flex items-center mb-[3%] border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black ">
                                <MdPermIdentity className='m-[1%] text-slate-700'/>
                                    <input
                                    type="text"
                                    name="name"
                                    placeholder="Full Name"
                                    value={formData.name}
                                    onChange={handleChange}
                                    className="ml-[2%] w-full"
                                    // className="w-full border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black"
                                    required
                                    />
                                </div> */}
                            </div>
                        </form>
                 </div>
            </div>
        </div>

  
      )

};

export default SignUpPage;