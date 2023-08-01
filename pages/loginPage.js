import React from 'react'
import { useState } from 'react';
import Image from 'next/image';
import Head from 'next/head';
import { FaRegEnvelope } from 'react-icons/fa';
import { FcGoogle } from 'react-icons/fc';
import {MdLockOutline} from 'react-icons/md'

export const loginPage = () => {
    
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
        console.log(formData);
        // Reset form setelah pengiriman
        setFormData({
          name: '',
          username: '',
          email: '',
          password: ''
        });
      };
    
      const handleGoogleRegistration = () => {
        // Lakukan proses pendaftaran dengan Google di sini
        console.log('Melakukan pendaftaran dengan Google');
      };
    return (
    <div className=" flex flex-col items-center justify-center min-h-screen py-2 w-full h-full bg-zinc-100">
        <div>
            <Head>
                <title>Sign In - Welcome to JICO</title>
            </Head>
        </div>
        <main className='flex flex-col items-center justify-center w-full flex-1 px-20 text-center'>
            <div className='container absolute rounded-2xl flex w-[80%] max-w-4xl shadow-[0_3px_10px_rgb(0,0,0,1)] bg-zinc-400'>



                <form onSubmit={handleSubmit} className='Background w-3/5 m-0 rounded-l-2xl bg-zinc-100 text-slate-700'>
                    
                    <div className=''>
                        <h1 className='text-slate-700 text-center text-4xl font-semibold leading-snug mt-[10%] '>Sign in</h1>
                        <div className='mt-[9%] flex flex-col items-center'>
                            <div className='w-[60%] bg-white shadow-[0_3px_10px_rgb(0,0,0,0.1)] rounded-md p-[3%] flex items-center mb-[3%]'>
                                <FaRegEnvelope className='m-[1%] text-slate-700'/>
                                <input className='ml-[2%] w-[90%]' type='email' name='email' value={formData.email} onChange={handleChange} placeholder='Email'/>
                            </div>
                            <div className='w-[60%] bg-white shadow-[0_3px_10px_rgb(0,0,0,0.1)] rounded-md p-[3%] flex items-center mb-[7%]'>
                                <MdLockOutline className='m-[1%] text-slate-700'/>
                                <input className='ml-[2%] w-[90%]' type='password' name='password' value={formData.password} onChange={handleChange} placeholder='Password'/>
                            </div>
                            <div className="text-center w-[60%] mb-[5%]">
                                <button type="submit" className="bg-blue-700 w-[100%] text-zinc-100 py-2 rounded-md hover:bg-blue-800 transition-colors duration-200">Sign In</button>
                            </div>
                            <div className="flex mb-[3%] items-center">
                                <label htmlFor="remember" className="flex items-center  text-justify text-xs">
                                    <input type="checkbox" name="remember" className="mr-2"
                                    />
                                    Remember me
                                </label>                   
                            </div>
                            <div className='flex flex-col justify-center  items-center w-[60%]'>
                                <p className='bg-zinc-100 text-slate-700 text-base font-normal leading-snug py-[2%] mt-[2%]'>or Connect With</p>
                                <button onClick={handleGoogleRegistration} className="w-[100%] bg-zinc-100 rounded-md border  hover:bg-gray-400 transition-colors duration-200 justify-center items-center flex py-[3%] mb-[2%] shadow-[0_3px_10px_rgb(0,0,0,0.1)]" ><FcGoogle size="25px" className="mr-[5%] "/> Sign In With Google</button>
                            </div>
                            <div className='ForgotPassword mt-[5%] '>
                                <a href='/forgotpass' className='text-slate-700 text-base font-normal leading-snug'>Forgot Password?</a>
                            </div>
                            <div className='mt-[1%] mb-[15%]'>
                                <p className='SignUp text-slate-700 text-base font-normal leading-snug'>Not a member yet? <a href='#' className='text-blue-700 font-bold text-base leading-snug'>Sign Up</a></p>
                            </div>
                        </div>
                    </div>
                </form>




                <div className='Background bg-gradient-to-t from-sky-200 to-stone-50 rounded-r-2xl m-0 ml-0 w-2/5 pt-28 pb-20 px-10'>
                    <h2 className='text-2xl font-semibold mb-2 text-right text-slate-700 leading-snug pt-[30%]'>Welcome to JICO</h2>
                    <p className='text-right text-slate-700 text-base font-normal leading-snug mb-20'>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam nec ultricies nisi. Suspendisse pulvinar viverra nibh vel ultricies. Mauris tincidunt mollis diam, at mollis enim aliquet eget. Fusce eros neque, pharetra eget tincidunt in, tincidunt nec tellus. </p>

                    <div className='h-8 flex object-right-bottom'>
                        <img className='' src='https://upload.wikimedia.org/wikipedia/id/thumb/c/c4/Telkom_Indonesia_2013.svg/1200px-Telkom_Indonesia_2013.svg.png'/>
                        <img className='' src='https://cdn.icon-icons.com/icons2/2699/PNG/512/atlassian_jira_logo_icon_170512.png'/>
                    </div>
                    
                </div>
            </div>
        </main> 
    </div>
  )
}
export default loginPage ;