import Link from "next/link";
import Header from "../components/header";
import {FaDiscord} from 'react-icons/fa'
import { useRouter } from "next/router";
import Content from "@/components/content";
import Features from "@/components/feature";
import Footer from "@/components/footer";
import Head from 'next/head';

export default function Home() {

  return (
    <div>
       <Head>
         <title>Home</title>
      </Head>
      <Header/>
      {/* <FaDiscord/> */}
      <Content/>
      <Features/>
      <Footer/>
    </div>
  )
}


