{-# LANGUAGE DeriveGeneric #-}

import Text.Pandoc
import Data.Aeson (FromJSON, ToJSON, decode, encode)
import qualified Data.ByteString.Lazy.Char8 as BL
import System.Environment
import System.ZMQ
import Data.ByteString.Char8
import Debug.Trace
import GHC.Generics (Generic)


data Input = Input{ htmlText :: String } 
             deriving (Show, Generic)

instance FromJSON Input
instance ToJSON Input


addr = "tcp://127.0.0.1:9999"

-- main :: IO ()
main =  withContext 64 serve


-- serve :: Context -> IO ()
serve context = withSocket context Rep process


-- process :: Socket a -> IO ()
process socket = do bind socket addr
                    Prelude.putStrLn "Accepting connections..."
                    handle socket


-- handle :: Socket a -> IO ()
handle socket = do readString socket >>= writeString socket . transform . getV . unmarshalJSON . toLazyBytes . unpack
                   handle socket


-- toLazyBytes :: [Char] -> ByteString
toLazyBytes s = BL.pack s


-- unmarshalJSON :: String -> String
-- unmarshalJSON x | trace ("unmarshal JSON from byteString: " ++ show x) False = undefined
unmarshalJSON bs = decode bs :: Maybe Input


-- getV :: Maybe Input -> String
-- getV a | trace ("vv" ++ show a) False = undefined
getV a = case a of
     Nothing -> "opppps!"
     Just x -> htmlText x


-- transform :: String -> String
transform = writeMarkdown def . readHtml def


-- writeString :: Socket a -> String -> IO ()
writeString socket string = send socket (pack string) []


-- readString :: Socket a -> IO ByteString
readString socket = receive socket []
