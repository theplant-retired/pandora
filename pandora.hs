{-# LANGUAGE OverloadedStrings #-}
{-# LANGUAGE DeriveGeneric #-}

import Control.Monad
import Control.Applicative ((<$>))
import Data.ByteString.Char8
import Data.ByteString.UTF8 (fromString)
import Data.Aeson (FromJSON, ToJSON, decode, encode)
import qualified Data.ByteString.Lazy.Char8 as BL
import Debug.Trace
import GHC.Generics (Generic)
import System.Environment
import System.ZMQ3.Monadic
import System.Posix.Daemon
import Text.Pandoc
import Text.Printf


data Input = Input{ 
                  text :: String ,
                  from :: String ,
                  to :: String
            } 
             deriving (Show, Generic)

instance FromJSON Input
instance ToJSON Input


addr = "tcp://127.0.0.1:9999"

workerURL = "inproc://workers"


main :: IO ()
main = runDetached (Nothing) def $ runZMQ $ do 
        -- Socket to talk to clients
        clients <- socket Router
        bind clients addr
    
        -- Socket to talk to workers
        workers <- socket Dealer
        bind workers workerURL
      
        -- using inproc (inter-thread) we expect to share the same context
        replicateM_ 1 (async worker)
        
        -- Connect work threads to client threads via a queue
        proxy clients workers Nothing


worker :: ZMQ z ()
worker = do
    receiver <- socket Rep
    connect receiver workerURL
    forever $ do
        unpack <$> receive receiver >>= send receiver [] . pack . utf8 . transform . unmarshalJSON . toLazyBytes


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
