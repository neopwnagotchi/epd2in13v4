package epd2in13v4

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

// Display resolution
const (
	Width  = 122
	Height = 250
)

var (
	PIN_RST  = rpio.Pin(17) // Reset
	PIN_DC   = rpio.Pin(25) // Data/Command
	PIN_CS   = rpio.Pin(8)  // SPI Chip Select (CE0)
	PIN_BUSY = rpio.Pin(24) // Busy
)

/******************************************************************************
function :	Software reset
parameter:
******************************************************************************/
/** static void EPD_2in13_V4_Reset(void)
{
    DEV_Digital_Write(EPD_RST_PIN, 1);
    DEV_Delay_ms(20);
    DEV_Digital_Write(EPD_RST_PIN, 0);
    DEV_Delay_ms(2);
    DEV_Digital_Write(EPD_RST_PIN, 1);
    DEV_Delay_ms(20);
}**/

func EPD_2in13_V4_Reset() {
	PIN_RST.High()
	time.Sleep(20 * time.Millisecond)
	PIN_RST.Low()
	time.Sleep(2 * time.Millisecond)
	PIN_RST.High()
	time.Sleep(20 * time.Millisecond)
}

/******************************************************************************
function :	send command
parameter:
     Reg : Command register
******************************************************************************/
/** static void EPD_2in13_V4_SendCommand(UBYTE Reg)
{
    DEV_Digital_Write(EPD_DC_PIN, 0);
    DEV_Digital_Write(EPD_CS_PIN, 0);
    DEV_SPI_WriteByte(Reg);
    DEV_Digital_Write(EPD_CS_PIN, 1);
}**/

func EPD_2in13_V4_SendCommand(Reg byte) {
	PIN_DC.Low()
	PIN_CS.Low()
	rpio.SpiTransmit(Reg)
	PIN_CS.High()
}

/******************************************************************************
function :	send data
parameter:
    Data : Write data
******************************************************************************/
/** static void EPD_2in13_V4_SendData(UBYTE Data)
{
    DEV_Digital_Write(EPD_DC_PIN, 1);
    DEV_Digital_Write(EPD_CS_PIN, 0);
    DEV_SPI_WriteByte(Data);
    DEV_Digital_Write(EPD_CS_PIN, 1);
}**/

func EPD_2in13_V4_SendData(Data byte) {
	PIN_DC.High()
	PIN_CS.Low()
	rpio.SpiTransmit(Data)
	PIN_CS.High()
}

/******************************************************************************
function :	Wait until the busy_pin goes LOW
parameter:
******************************************************************************/
/** void EPD_2in13_V4_ReadBusy(void)
{
    Debug("e-Paper busy\r\n");
	while(1)
	{	 //=1 BUSY
		if(DEV_Digital_Read(EPD_BUSY_PIN)==0)
			break;
		DEV_Delay_ms(10);
	}
	DEV_Delay_ms(10);
    Debug("e-Paper busy release\r\n");
}**/

func EPD_2in13_V4_ReadBusy() {
	fmt.Println("e-Paper busy")
	for {
		if PIN_BUSY.Read() == 0 {
			break
		}
		time.Sleep(10 * time.Microsecond)
	}
	time.Sleep(10 * time.Microsecond)
	fmt.Println("e-Paper busy release")
}

/******************************************************************************
function :	Setting the display window
parameter:
	Xstart : X-axis starting position
	Ystart : Y-axis starting position
	Xend : End position of X-axis
	Yend : End position of Y-axis
******************************************************************************/
/** static void EPD_2in13_V4_SetWindows(UWORD Xstart, UWORD Ystart, UWORD Xend, UWORD Yend)
{
    EPD_2in13_V4_SendCommand(0x44); // SET_RAM_X_ADDRESS_START_END_POSITION
    EPD_2in13_V4_SendData((Xstart>>3) & 0xFF);
    EPD_2in13_V4_SendData((Xend>>3) & 0xFF);

    EPD_2in13_V4_SendCommand(0x45); // SET_RAM_Y_ADDRESS_START_END_POSITION
    EPD_2in13_V4_SendData(Ystart & 0xFF);
    EPD_2in13_V4_SendData((Ystart >> 8) & 0xFF);
    EPD_2in13_V4_SendData(Yend & 0xFF);
    EPD_2in13_V4_SendData((Yend >> 8) & 0xFF);
}**/

func EPD_2in13_V4_SetWindows(Xstart, Ystart, Xend, Yend int) {
	EPD_2in13_V4_SendCommand(0x44) // SET_RAM_X_ADDRESS_START_END_POSITION
	EPD_2in13_V4_SendData(byte((Xstart >> 3) & 0xFF))
	EPD_2in13_V4_SendData(byte((Xend >> 3) & 0xFF))

	EPD_2in13_V4_SendCommand(0x45) // SET_RAM_Y_ADDRESS_START_END_POSITION
	EPD_2in13_V4_SendData(byte(Ystart & 0xFF))
	EPD_2in13_V4_SendData(byte((Ystart >> 8) & 0xFF))
	EPD_2in13_V4_SendData(byte(Yend & 0xFF))
	EPD_2in13_V4_SendData(byte((Yend >> 8) & 0xFF))
}

/******************************************************************************
function :	Set Cursor
parameter:
	Xstart : X-axis starting position
	Ystart : Y-axis starting position
******************************************************************************/
/** static void EPD_2in13_V4_SetCursor(UWORD Xstart, UWORD Ystart)
{
    EPD_2in13_V4_SendCommand(0x4E); // SET_RAM_X_ADDRESS_COUNTER
    EPD_2in13_V4_SendData(Xstart & 0xFF);

    EPD_2in13_V4_SendCommand(0x4F); // SET_RAM_Y_ADDRESS_COUNTER
    EPD_2in13_V4_SendData(Ystart & 0xFF);
    EPD_2in13_V4_SendData((Ystart >> 8) & 0xFF);
}**/

func EPD_2in13_V4_SetCursor(Xstart, Ystart int) {
	EPD_2in13_V4_SendCommand(0x4E) // SET_RAM_X_ADDRESS_COUNTER
	EPD_2in13_V4_SendData(byte(Xstart & 0xFF))

	EPD_2in13_V4_SendCommand(0x4F) // SET_RAM_Y_ADDRESS_COUNTER
	EPD_2in13_V4_SendData(byte(Ystart & 0xFF))
	EPD_2in13_V4_SendData(byte((Ystart >> 8) & 0xFF))
}

/******************************************************************************
function :	Turn On Display
parameter:
******************************************************************************/
/** static void EPD_2in13_V4_TurnOnDisplay(void)
{
	EPD_2in13_V4_SendCommand(0x22); // Display Update Control
	EPD_2in13_V4_SendData(0xf7);
	EPD_2in13_V4_SendCommand(0x20); // Activate Display Update Sequence
	EPD_2in13_V4_ReadBusy();
}**/

func EPD_2in13_V4_TurnOnDisplay() {
	EPD_2in13_V4_SendCommand(0x22) // Display Update Control
	EPD_2in13_V4_SendData(0xf7)
	EPD_2in13_V4_SendCommand(0x20) // Activate Display Update Sequence
	EPD_2in13_V4_ReadBusy()
}

/** static void EPD_2in13_V4_TurnOnDisplay_Fast(void)
{
	EPD_2in13_V4_SendCommand(0x22); // Display Update Control
	EPD_2in13_V4_SendData(0xc7);	// fast:0x0c, quality:0x0f, 0xcf
	EPD_2in13_V4_SendCommand(0x20); // Activate Display Update Sequence
	EPD_2in13_V4_ReadBusy();
}**/

func EPD_2in13_V4_TurnOnDisplay_Fast() {
	EPD_2in13_V4_SendCommand(0x22) // Display Update Control
	EPD_2in13_V4_SendData(0xc7)    // fast:0x0c, quality:0x0f, 0xcf
	EPD_2in13_V4_SendCommand(0x20) // Activate Display Update Sequence
	EPD_2in13_V4_ReadBusy()
}

/** static void EPD_2in13_V4_TurnOnDisplay_Partial(void)
{
	EPD_2in13_V4_SendCommand(0x22); // Display Update Control
	EPD_2in13_V4_SendData(0xff);	// fast:0x0c, quality:0x0f, 0xcf
	EPD_2in13_V4_SendCommand(0x20); // Activate Display Update Sequence
	EPD_2in13_V4_ReadBusy();
}**/

func EPD_2in13_V4_TurnOnDisplay_Partial() {
	EPD_2in13_V4_SendCommand(0x22) // Display Update Control
	EPD_2in13_V4_SendData(0xff)    // fast:0x0c, quality:0x0f, 0xcf
	EPD_2in13_V4_SendCommand(0x20) // Activate Display Update Sequence
	EPD_2in13_V4_ReadBusy()
}

/******************************************************************************
function :	Initialize the e-Paper register
parameter:
******************************************************************************/
/** void EPD_2in13_V4_Init(void)
{
	EPD_2in13_V4_Reset();

	EPD_2in13_V4_ReadBusy();
	EPD_2in13_V4_SendCommand(0x12);  //SWRESET
	EPD_2in13_V4_ReadBusy();

	EPD_2in13_V4_SendCommand(0x01); //Driver output control
	EPD_2in13_V4_SendData(0xF9);
	EPD_2in13_V4_SendData(0x00);
	EPD_2in13_V4_SendData(0x00);

	EPD_2in13_V4_SendCommand(0x11); //data entry mode
	EPD_2in13_V4_SendData(0x03);

	EPD_2in13_V4_SetWindows(0, 0, EPD_2in13_V4_WIDTH-1, EPD_2in13_V4_HEIGHT-1);
	EPD_2in13_V4_SetCursor(0, 0);

	EPD_2in13_V4_SendCommand(0x3C); //BorderWavefrom
	EPD_2in13_V4_SendData(0x05);

	EPD_2in13_V4_SendCommand(0x21); //  Display update control
	EPD_2in13_V4_SendData(0x00);
	EPD_2in13_V4_SendData(0x80);

	EPD_2in13_V4_SendCommand(0x18); //Read built-in temperature sensor
	EPD_2in13_V4_SendData(0x80);
	EPD_2in13_V4_ReadBusy();
}**/

func EPD_2in13_V4_Init() error {

	if err := rpio.Open(); err != nil {
		return fmt.Errorf("failed to open GPIO: %v", err)
	}

	rpio.SpiBegin(rpio.Spi0)
	rpio.SpiMode(0, 0)
	rpio.SpiSpeed(4000000)

	PIN_RST.Output()
	PIN_DC.Output()
	PIN_CS.Output()
	PIN_BUSY.Input()

	EPD_2in13_V4_Reset()

	EPD_2in13_V4_ReadBusy()
	EPD_2in13_V4_SendCommand(0x12) // SWRESET
	EPD_2in13_V4_ReadBusy()

	EPD_2in13_V4_SendCommand(0x01) // Driver output control
	EPD_2in13_V4_SendData(0xF9)
	EPD_2in13_V4_SendData(0x00)
	EPD_2in13_V4_SendData(0x00)

	EPD_2in13_V4_SendCommand(0x11) // data entry mode
	EPD_2in13_V4_SendData(0x03)

	EPD_2in13_V4_SetWindows(0, 0, Width-1, Height-1)
	EPD_2in13_V4_SetCursor(0, 0)

	EPD_2in13_V4_SendCommand(0x3C)
	EPD_2in13_V4_SendData(0x05)

	EPD_2in13_V4_SendCommand(0x21) // Display update control
	EPD_2in13_V4_SendData(0x00)
	EPD_2in13_V4_SendData(0x80)

	EPD_2in13_V4_SendCommand(0x18) // Read built-in temperature sensor
	EPD_2in13_V4_SendData(0x80)
	EPD_2in13_V4_ReadBusy()

	return nil
}

/** void EPD_2in13_V4_Init_Fast(void)
{
	EPD_2in13_V4_Reset();

	EPD_2in13_V4_SendCommand(0x12);  //SWRESET
	EPD_2in13_V4_ReadBusy();

	EPD_2in13_V4_SendCommand(0x18); //Read built-in temperature sensor
	EPD_2in13_V4_SendData(0x80);

	EPD_2in13_V4_SendCommand(0x11); //data entry mode
	EPD_2in13_V4_SendData(0x03);

	EPD_2in13_V4_SetWindows(0, 0, EPD_2in13_V4_WIDTH-1, EPD_2in13_V4_HEIGHT-1);
	EPD_2in13_V4_SetCursor(0, 0);

	EPD_2in13_V4_SendCommand(0x22); // Load temperature value
	EPD_2in13_V4_SendData(0xB1);
	EPD_2in13_V4_SendCommand(0x20);
	EPD_2in13_V4_ReadBusy();

	EPD_2in13_V4_SendCommand(0x1A); // Write to temperature register
	EPD_2in13_V4_SendData(0x64);
	EPD_2in13_V4_SendData(0x00);

	EPD_2in13_V4_SendCommand(0x22); // Load temperature value
	EPD_2in13_V4_SendData(0x91);
	EPD_2in13_V4_SendCommand(0x20);
	EPD_2in13_V4_ReadBusy();
}**/

func EPD_2in13_V4_Init_Fast() error {

	if err := rpio.Open(); err != nil {
		return fmt.Errorf("failed to open GPIO: %v", err)
	}

	rpio.SpiBegin(rpio.Spi0)
	rpio.SpiMode(0, 0)
	rpio.SpiSpeed(4000000)

	PIN_RST.Output()
	PIN_DC.Output()
	PIN_CS.Output()
	PIN_BUSY.Input()

	EPD_2in13_V4_Reset()

	EPD_2in13_V4_SendCommand(0x12) // SWRESET
	EPD_2in13_V4_ReadBusy()

	EPD_2in13_V4_SendCommand(0x18) // Read built-in temperature sensor
	EPD_2in13_V4_SendData(0x80)

	EPD_2in13_V4_SendCommand(0x11) // data entry mode
	EPD_2in13_V4_SendData(0x03)

	EPD_2in13_V4_SetWindows(0, 0, Width-1, Height-1)
	EPD_2in13_V4_SetCursor(0, 0)

	EPD_2in13_V4_SendCommand(0x22) // Load temperature value
	EPD_2in13_V4_SendData(0xB1)
	EPD_2in13_V4_SendCommand(0x20)
	EPD_2in13_V4_ReadBusy()

	EPD_2in13_V4_SendCommand(0x1A) // Write to temperature register
	EPD_2in13_V4_SendData(0x64)
	EPD_2in13_V4_SendData(0x00)

	EPD_2in13_V4_SendCommand(0x22) // Load temperature value
	EPD_2in13_V4_SendData(0x91)
	EPD_2in13_V4_SendCommand(0x20)
	EPD_2in13_V4_ReadBusy()

	return nil
}

/******************************************************************************
function :	Clear screen
parameter:
******************************************************************************/
/** void EPD_2in13_V4_Clear(void)
{
	UWORD Width, Height;
    Width = (EPD_2in13_V4_WIDTH % 8 == 0)? (EPD_2in13_V4_WIDTH / 8 ): (EPD_2in13_V4_WIDTH / 8 + 1);
    Height = EPD_2in13_V4_HEIGHT;

    EPD_2in13_V4_SendCommand(0x24);
    for (UWORD j = 0; j < Height; j++) {
        for (UWORD i = 0; i < Width; i++) {
            EPD_2in13_V4_SendData(0XFF);
        }
    }

	EPD_2in13_V4_TurnOnDisplay();
}**/

func EPD_2in13_V4_Clear() {
	width := (Width % 8)
	if width == 0 {
		width = Width / 8
	} else {
		width = (Width / 8) + 1
	}
	height := (Height)

	EPD_2in13_V4_SendCommand(0x24)
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			EPD_2in13_V4_SendData(0xFF)
		}
	}

	EPD_2in13_V4_TurnOnDisplay()
}

/** void EPD_2in13_V4_Clear_Black(void)
{
	UWORD Width, Height;
    Width = (EPD_2in13_V4_WIDTH % 8 == 0)? (EPD_2in13_V4_WIDTH / 8 ): (EPD_2in13_V4_WIDTH / 8 + 1);
    Height = EPD_2in13_V4_HEIGHT;

    EPD_2in13_V4_SendCommand(0x24);
    for (UWORD j = 0; j < Height; j++) {
        for (UWORD i = 0; i < Width; i++) {
            EPD_2in13_V4_SendData(0X00);
        }
    }

	EPD_2in13_V4_TurnOnDisplay();
}**/

func EPD_2in13_V4_Clear_Black() {
	width := (Width % 8)
	if width == 0 {
		width = Width / 8
	} else {
		width = (Width / 8) + 1
	}
	height := (Height)

	EPD_2in13_V4_SendCommand(0x24)
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			EPD_2in13_V4_SendData(0x00)
		}
	}

	EPD_2in13_V4_TurnOnDisplay()
}

/******************************************************************************
function :	Sends the image buffer in RAM to e-Paper and displays
parameter:
	Image : Image data
******************************************************************************/
/** void EPD_2in13_V4_Display(UBYTE *Image)
{
	UWORD Width, Height;
    Width = (EPD_2in13_V4_WIDTH % 8 == 0)? (EPD_2in13_V4_WIDTH / 8 ): (EPD_2in13_V4_WIDTH / 8 + 1);
    Height = EPD_2in13_V4_HEIGHT;

    EPD_2in13_V4_SendCommand(0x24);
    for (UWORD j = 0; j < Height; j++) {
        for (UWORD i = 0; i < Width; i++) {
            EPD_2in13_V4_SendData(Image[i + j * Width]);
        }
    }

	EPD_2in13_V4_TurnOnDisplay();
}**/

func EPD_2in13_V4_Display(Image []byte) {
	width := (Width % 8)
	if width == 0 {
		width = Width / 8
	} else {
		width = (Width / 8) + 1
	}
	height := (Height)
	EPD_2in13_V4_SendCommand(0x24)
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			EPD_2in13_V4_SendData(Image[i+j*width])
		}
	}

	EPD_2in13_V4_TurnOnDisplay()
}

/** void EPD_2in13_V4_Display_Fast(UBYTE *Image)
{
	UWORD Width, Height;
    Width = (EPD_2in13_V4_WIDTH % 8 == 0)? (EPD_2in13_V4_WIDTH / 8 ): (EPD_2in13_V4_WIDTH / 8 + 1);
    Height = EPD_2in13_V4_HEIGHT;

    EPD_2in13_V4_SendCommand(0x24);
    for (UWORD j = 0; j < Height; j++) {
        for (UWORD i = 0; i < Width; i++) {
            EPD_2in13_V4_SendData(Image[i + j * Width]);
        }
    }

	EPD_2in13_V4_TurnOnDisplay_Fast();
}**/

func EPD_2in13_V4_Display_Fast(Image []byte) {
	width := (Width % 8)
	if width == 0 {
		width = Width / 8
	} else {
		width = (Width / 8) + 1
	}
	height := (Height)
	EPD_2in13_V4_SendCommand(0x24)
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			EPD_2in13_V4_SendData(Image[i+j*width])
		}
	}

	EPD_2in13_V4_TurnOnDisplay_Fast()

}

/******************************************************************************
function :	Refresh a base image
parameter:
	Image : Image data
******************************************************************************/
/**void EPD_2in13_V4_Display_Base(UBYTE *Image)
{
	UWORD Width, Height;
    Width = (EPD_2in13_V4_WIDTH % 8 == 0)? (EPD_2in13_V4_WIDTH / 8 ): (EPD_2in13_V4_WIDTH / 8 + 1);
    Height = EPD_2in13_V4_HEIGHT;

	EPD_2in13_V4_SendCommand(0x24);   //Write Black and White image to RAM
    for (UWORD j = 0; j < Height; j++) {
        for (UWORD i = 0; i < Width; i++) {
			EPD_2in13_V4_SendData(Image[i + j * Width]);
		}
	}
	EPD_2in13_V4_SendCommand(0x26);   //Write Black and White image to RAM
    for (UWORD j = 0; j < Height; j++) {
        for (UWORD i = 0; i < Width; i++) {
			EPD_2in13_V4_SendData(Image[i + j * Width]);
		}
	}
	EPD_2in13_V4_TurnOnDisplay();
}**/

func EPD_2in13_V4_Display_Base(Image []byte) {
	width := (Width % 8)
	if width == 0 {
		width = Width / 8
	} else {
		width = (Width / 8) + 1
	}
	height := (Height)
	EPD_2in13_V4_SendCommand(0x24) // Write Black and White image to RAM
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			EPD_2in13_V4_SendData(Image[i+j*width])
		}
	}
	EPD_2in13_V4_SendCommand(0x26) // Write Black and White image to RAM
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			EPD_2in13_V4_SendData(Image[i+j*width])
		}
	}
	EPD_2in13_V4_TurnOnDisplay()
}

/******************************************************************************
function :	Sends the image buffer in RAM to e-Paper and partial refresh
parameter:
	Image : Image data
******************************************************************************/
/**void EPD_2in13_V4_Display_Partial(UBYTE *Image)
{
	UWORD Width, Height;
    Width = (EPD_2in13_V4_WIDTH % 8 == 0)? (EPD_2in13_V4_WIDTH / 8 ): (EPD_2in13_V4_WIDTH / 8 + 1);
    Height = EPD_2in13_V4_HEIGHT;

	//Reset
    DEV_Digital_Write(EPD_RST_PIN, 0);
    DEV_Delay_ms(1);
    DEV_Digital_Write(EPD_RST_PIN, 1);

	EPD_2in13_V4_SendCommand(0x3C); //BorderWavefrom
	EPD_2in13_V4_SendData(0x80);

	EPD_2in13_V4_SendCommand(0x01); //Driver output control
	EPD_2in13_V4_SendData(0xF9);
	EPD_2in13_V4_SendData(0x00);
	EPD_2in13_V4_SendData(0x00);

	EPD_2in13_V4_SendCommand(0x11); //data entry mode
	EPD_2in13_V4_SendData(0x03);

	EPD_2in13_V4_SetWindows(0, 0, EPD_2in13_V4_WIDTH-1, EPD_2in13_V4_HEIGHT-1);
	EPD_2in13_V4_SetCursor(0, 0);

	EPD_2in13_V4_SendCommand(0x24);   //Write Black and White image to RAM
    for (UWORD j = 0; j < Height; j++) {
        for (UWORD i = 0; i < Width; i++) {
			EPD_2in13_V4_SendData(Image[i + j * Width]);
		}
	}
	EPD_2in13_V4_TurnOnDisplay_Partial();
}**/

func EPD_2in13_V4_Display_Partial(Image []byte) {
	width := (Width % 8)
	if width == 0 {
		width = Width / 8
	} else {
		width = (Width / 8) + 1
	}
	height := (Height)

	PIN_RST.Low()
	time.Sleep(1 * time.Millisecond)
	PIN_RST.High()

	EPD_2in13_V4_SendCommand(0x3C) // BorderWavefrom
	EPD_2in13_V4_SendData(0x80)

	EPD_2in13_V4_SendCommand(0x01) // Driver output control
	EPD_2in13_V4_SendData(0xF9)
	EPD_2in13_V4_SendData(0x00)
	EPD_2in13_V4_SendData(0x00)

	EPD_2in13_V4_SendCommand(0x11) // data entry mode
	EPD_2in13_V4_SendData(0x03)

	EPD_2in13_V4_SetWindows(0, 0, Width-1, Height-1)
	EPD_2in13_V4_SetCursor(0, 0)

	EPD_2in13_V4_SendCommand(0x24) // Write Black and White image to RAM
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			EPD_2in13_V4_SendData(Image[i+j*width])
		}
	}

	EPD_2in13_V4_TurnOnDisplay_Partial()
}

/******************************************************************************
function :	Enter sleep mode
parameter:
******************************************************************************/
/** void EPD_2in13_V4_Sleep(void)
{
	EPD_2in13_V4_SendCommand(0x10); //enter deep sleep
	EPD_2in13_V4_SendData(0x01);
	DEV_Delay_ms(100);
} **/

func EPD_2in13_V4_Sleep() {
	EPD_2in13_V4_SendCommand(0x10) // enter deep sleep
	EPD_2in13_V4_SendData(0x01)
	time.Sleep(100 * time.Millisecond)
}
