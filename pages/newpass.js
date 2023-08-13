import { useState } from "react";
import Header from "@/components/header";
import { useRouter } from "next/router";
import Head from 'next/head';
import {MdLockOutline} from 'react-icons/md'

export default function Forgotpass() {
    const [formData, setFormData] = useState({
        newPassword: '',
        confirmPassword: ''
      });
    
      const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
      };
    
      const handleSubmit = (e) => {
        e.preventDefault();
        // Lakukan aksi reset password di sini
        console.log(formData);
        // Reset form setelah pengiriman
        setFormData({
            newPassword: '',
            confirmPassword: ''
        });
      };

      const router = useRouter();

      const handleButtonClick = () => {
        router.push('/loginPage');
      };
    
      return (
        // <div className="flex flex-col items-center justify-center min-h-screen py-2 bg-zinc-100">
        //    <div>
        //     <Head>
        //       <title>Forgot Password</title>
        //     </Head>
        //   </div>
        //   <main className="flex flex-col items-center justify-center w-full flex-1 px-20 text-center">
        //     <form onSubmit={handleSubmit} className="w-full max-w-md bg-white rounded-lg shadow-md p-8">
        //       <h2 className="text-2xl font-semibold mb-4">Forgot Password</h2>
        //       <div className="mb-4">
        //         <input
        //           className="w-full px-3 py-2 placeholder-gray-400 border rounded-lg focus:outline-none"
        //           type="email"
        //           name="email"
        //           value={formData.email}
        //           onChange={handleChange}
        //           required
        //           placeholder="Email"
        //         />
        //       </div>
        //       <div className="mb-4">
        //         <input
        //           className="w-full px-3 py-2 placeholder-gray-400 border rounded-lg focus:outline-none"
        //           type="password"
        //           name="newPassword"
        //           value={formData.newPassword}
        //           onChange={handleChange}
        //           required
        //           placeholder="New Password"
        //         />
        //       </div>
        //       <div className="mb-4">
        //         <input
        //           className="w-full px-3 py-2 placeholder-gray-400 border rounded-lg focus:outline-none"
        //           type="password"
        //           name="confirmPassword"
        //           value={formData.confirmPassword}
        //           onChange={handleChange}
        //           required
        //           placeholder="Confirm Password"
        //         />
        //       </div>
        //       <div className="flex items-center justify-center">
        //         <button
        //           type="submit"
        //           className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-lg focus:outline-none"
        //           onClick={handleButtonClick}
        //         >
        //           Reset Password
        //         </button>
        //       </div>
        //     </form>
        //   </main>
        // </div>
        <div className=" flex flex-col items-center justify-center min-h-screen py-2 w-full h-full bg-zinc-100">       
        <div>
          <Head>
            <title>New Password</title>
          </Head>
        </div>

            {/* Left side */}
            <main className='flex flex-col items-center justify-center w-full flex-1 px-20 text-center'>
                <div className='container absolute rounded-2xl flex w-[80%] h-5/6 max-w-4xl shadow-[0_3px_10px_rgb(0,0,0,1)] bg-zinc-400'>

                    <form onSubmit={handleSubmit} className='Background w-3/5 m-0 rounded-l-2xl bg-zinc-100 text-slate-700'>
                        
                        <div className=''>
                            <h1 className='text-slate-700 text-center text-4xl font-semibold leading-snug mt-[15%] mb-3 '>Reset Password</h1>
                            <div className='mt-[9%] flex flex-col items-center'>
                                {/* New Password */}
                                <div className='w-[60%] bg-white shadow-[0_3px_10px_rgb(0,0,0,0.1)] rounded-md p-[3%] flex items-center mb-[3%]'>
                                    <MdLockOutline className='m-[1%] text-slate-700'/>
                                    <input className='ml-[2%] w-[90%]' type='password' name='password' value={formData.newPassword} onChange={handleChange} placeholder='New Password'/>
                                 </div>
                                
                                {/* Confirm Password */}
                                <div className='w-[60%] bg-white shadow-[0_3px_10px_rgb(0,0,0,0.1)] rounded-md p-[3%] flex items-center mb-[7%]'>
                                    <MdLockOutline className='m-[1%] text-slate-700'/>
                                    <input className='ml-[2%] w-[90%]' type='password' name='password' value={formData.confirmPassword} onChange={handleChange} placeholder='Confirm New Password'/>
                                 </div>
                            
                                <div className="flex items-center justify-center w-[60%] mb-[5%]">
                                <button
                                    type="submit"
                                    className="bg-blue-700 w-[100%] text-zinc-100 py-2 rounded-md hover:bg-blue-800 transition-colors duration-200"
                                    onClick={handleButtonClick}
                                >
                                    Reset Password
                                </button>
                                </div>                                
                            </div>
                        </div>
                    </form>


                {/* Right Side */}
                    <div className='Background bg-gradient-to-t from-sky-200 to-stone-50 rounded-r-2xl m-0 ml-0 w-2/5 pt-28 pb-20 px-10 text-right'>
                        <h2 className='text-3xl font-bold mb-2 text-slate-700 leading-snug '>Forgot Password?  </h2>
                        <p className=' text-slate-700 text-base font-normal leading-snug '>Set your new Password that you will remember!! </p>
                        <p className=' text-slate-700 text-base leading-snug mb-40 font-semibold'>Note : It must be different from your previous password </p>

                        <div className='h-8 flex object-right-bottom'>
                            <img className='' src='https://upload.wikimedia.org/wikipedia/id/thumb/c/c4/Telkom_Indonesia_2013.svg/1200px-Telkom_Indonesia_2013.svg.png'/>
                            <img className='' src='https://cdn.icon-icons.com/icons2/2699/PNG/512/atlassian_jira_logo_icon_170512.png'/>
                        </div>
                        
                    </div>
                </div>
            </main> 
        </div>

        );
    }

    