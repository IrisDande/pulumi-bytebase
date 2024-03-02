using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Bytebase = Pulumi.Bytebase;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new Bytebase.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

