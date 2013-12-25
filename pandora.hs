{-# LANGUAGE DeriveGeneric #-}

import Control.Monad
import Data.ByteString.Char8
import Data.ByteString.UTF8 (fromString)
import Data.Aeson (FromJSON, ToJSON, decode, encode)
import qualified Data.ByteString.Lazy.Char8 as BL
import Debug.Trace
import GHC.Generics (Generic)
import System.Environment
import System.ZMQ
import System.Posix.Daemon
import Text.Pandoc


data Input = Input{ 
                  text :: String ,
                  from :: String ,
                  to :: String
            } 
             deriving (Show, Generic)

instance FromJSON Input
instance ToJSON Input


addr = "tcp://127.0.0.1:9999"


main :: IO ()
main =  runDetached (Just "panda.pid") def $ forever $ do
-- main =  do  -- for Debug
  withContext 64 serve


serve :: Context -> IO ()
serve context = withSocket context Rep process


process :: Socket a -> IO ()
process socket = do bind socket addr
                    Prelude.putStrLn "Accepting connections..."
                    handle socket


handle :: Socket a -> IO ()
handle socket = do readString socket >>= writeString socket . utf8 . transform . unmarshalJSON . toLazyBytes . unpack
                   handle socket


toLazyBytes :: [Char] -> BL.ByteString
toLazyBytes s = BL.pack s


unmarshalJSON :: BL.ByteString -> Maybe Input
unmarshalJSON bs = decode bs :: Maybe Input


transform :: Maybe Input -> String
-- transform a | trace ("transform: " ++ show a) False = undefined
transform a = case a of
    Nothing -> "opppps!"
    Just x -> do 
        let f = from x in case f of
         "html" -> let t = to x in case t of
              "markdown" -> html2markdown $ text x
              "html" ->  text x
         "markdown" -> let t = to x in case t of
              "html" -> markdown2html $ text x
              "markdown" -> text x


utf8 :: String -> String 
utf8 = unpack . fromString


html2markdown :: String -> String
-- html2markdown a | trace ("html: " ++ show a) False = undefined
html2markdown = writeMarkdown def . readHtml def


markdown2html :: String -> String
markdown2html = writeHtmlString def . readMarkdown def


writeString :: Socket a -> String -> IO ()
writeString socket string = send socket (pack string) []


readString :: Socket a -> IO ByteString
readString socket = receive socket []
