import React from 'react'
import Head from 'next/head';
import { useState, useRef, useEffect } from 'react';
import { FaTimes } from 'react-icons/fa';
import { MdPermIdentity, MdLockOutline } from 'react-icons/md';
import { BiSolidHide, BiSolidInfoCircle } from "react-icons/bi";
import { FcCheckmark} from "react-icons/fc";
import axios from './api/axios'


const SigninPage = () => {
    const userReference = useRef();
    const errReference = useRef(); 

    const [uname, setUname] = useState('');    
    const [pass, setPass] = useState('');

    // Error Message
    const [errMsg, setErrMsg] = useState('');
    const [success, setSuccess] = useState(false);

    useEffect (()=> {
      userReference.current.focus(); //setting focusnya ketika komponen load
  }, []) 

  useEffect (() => {
    setErrMsg('');
  }, [uname, pass])

    // useEffect (() => {
    //   const result = username_valid.test(uname);
    //   console.log(result);
    //   console.log(uname);
    //   setValidUname(result);
    // }, [uname])

    // useEffect (() => {
    //   const result = pass_valid.test(pass);
    //   console.log(result);
    //   console.log(pass);
    //   setValidPass(result);
    // }, [pass])
  
    
      const handleSubmit = async (e) => {
        e.preventDefault();

        // try {
        //   const resp = await axios.post(signin_url,
        //     JSON.stringify({uname, pass}),{
        //       headers: {'Content-Type': 'application/json'},
        //       withCredentials: true
        //     }
        //   )
        //   console.log(resp.data);
        //   // console.log(resp.accesstoken)
        //   console.log(JSON.stringify(resp));
        //   setSuccess(true);

        // } catch(err){
        //   if (!err.resp){
        //     setErrMsg('No Server Response')
        //   }else{
        //     setErrMsg('Registration Failed')
        //   }
        //   errReference.current.focus();
        // }

        console.log( uname, pass)
        console.log('Pendaftaran berhasil');
        setUname('')
        setPass('')
        // console.log(resp.accesstoken)
        setSuccess(true);
        // Reset form setelah pengiriman
        // setFormData({
        //   name: '',
        //   username: '',
        //   email: '',
        //   password: ''
        // });
      };
  
    return (
    <div>
        <div>
            <Head>
                <title>Sign In - Welcome to JICO</title>
            </Head>
        </div>
        {/* background */}
        <div className='bg-gradient-to-r from-cyan-500/10 via-teal-300/10 to-sky-200/10 block h-screen items-center justify-center p-4 md:flex'>
            {/* container */}
            <div className='bg-gradient-to-t from-sky-200 to-stone-50 flex flex-col items-center max-w-screen-lg overflow-hidden rounded-lg shadow-[0_3px_10px_rgb(0,0,0,1)] w-full md:flex-row text-white'>
                {/* form card */}
                <div className='bg-white flex flex-col items-center p-4 space-y-8 w-full md:w-3/5 text-slate-700'>
                  
                    {/* TITLE Sign In */}
                    <div className='flex flex-col items-center'>
                        <h1 className=' text-slate-700 text-4xl font-semibold pt-14'>Sign in</h1>
                    </div>

                    <p ref={errReference} className={` ${errMsg ? "errmsg" : "offscreen"}`} aria-live="assertive">{errMsg}</p>

                    {/* ISI FORM */}
                    <form onSubmit={handleSubmit} className='flex flex-col items-center space-y-4'>

                       {/*Username  */}
                      <div className="w-full bg-white flex items-center mb-[3%] border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black ">
                      <MdPermIdentity className='m-[1%] text-slate-700'/>
                      <input className="pl-2 py-1 w-72 focus:outline-none"
                        type="text"
                        name="username"
                        placeholder="Username"
                        ref={userReference}
                        autoComplete="off"
                        onChange={(e) => setUname(e.target.value)}                       
                        required
                        value = {uname}
                      /> 
                      </div>

                        {/* Password */}
                            <div className=' w-full bg-white flex items-center mb-[3%] border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black'>
                                <MdLockOutline className='m-[1%] text-slate-700'/>
                                <input className="pl-2 py-1 w-72 focus:outline-none"
                                   type="password"
                                   name="password"
                                   placeholder="Password"
                                   onChange={(e) => setPass(e.target.value)}
                                   value = {pass}
                                   required
                                />
                                 <BiSolidHide/>
                            </div>

                        {/* Submit */}
                        <div className='relative py-2'>
                            <button href="/homePage" type="submit" className=" bg-blue-700  text-zinc-100 text-center justify-center rounded-md hover:bg-blue-800 transition-colors duration-200 inline-flex items-center w-72 h-10 px-3 py-1">Sign In</button>
                        </div>
                    </form>

                    <div className='flex flex-col items-center'>
                        <label htmlFor="remember" className="flex items-center  text-justify text-xs">
                            <input type="checkbox" name="remember" className="mr-2"/>
                                Remember me
                        </label>
                    </div>

                    <div className='flex flex-col items-center pb-10'>
                        <a href='/ForgotPass' className='text-slate-700 text-s font-semibold leading-snug hover:underline'>Forgot Password?</a>

                        <p className='SignUp text-slate-700 text-base font-normal leading-snug'>Not a member yet? <a href='/signUpPage' className='text-blue-700 font-bold text-base leading-snug hover:underline'>Sign Up</a></p>
                    </div>
                    <div className=''>
                        
                    </div>

                </div>

                {/* welcome card */}
                <div className='flex flex-col items-center justify-center text-right text-slate-700 p-4 w-full md:w-2/5'>
                    <div>

                        <h2 className='text-2xl font-bold mb-2'>Welcome to JICO</h2>
                        <p className='text-base font-normal mb-48'>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam nec ultricies nisi. Suspendisse pulvinar viverra nibh vel ultricies. Mauris tincidunt mollis diam, at mollis enim aliquet eget. Fusce eros neque, pharetra eget tincidunt in, tincidunt nec tellus. </p>

                        <div className='h-8 flex object-right-bottom'>
                        <img className='' src='https://upload.wikimedia.org/wikipedia/id/thumb/c/c4/Telkom_Indonesia_2013.svg/1200px-Telkom_Indonesia_2013.svg.png'/>
                        <img className='' src='https://cdn.icon-icons.com/icons2/2699/PNG/512/atlassian_jira_logo_icon_170512.png'/>
                        </div>
                    </div>
                </div>

                  {/* pop up Success */}
                  {success && (
                        <div className='fixed inset-0 flex items-center justify-center z-50 bg-black opacity-80'>
                          <div className='bg-slate-900 p-8 rounded-lg flex flex-col text-white'>
                            <h2 className='text-xl font-semibold mb-4 text-white text-center'>Successfully Sign In</h2>
                            <FcCheckmark className="text-6xl text-emerald-500 animate-pulse mx-auto mb-4"/>
                            <a href= '/' className="text-center  hover:underline ">Homepage</a>
                            <button
                              onClick={() => setSuccess(false)} 
                              className='bg-emerald-600 px-4 py-2 hover:bg-emerald-400 font-bold rounded-lg mt-4 text-white'
                            >
                              Close
                            </button> 
                          </div>
                        </div>
                        )}
                
            </div>        
        </div>
    </div>
    
  )
}
export default SigninPage ;